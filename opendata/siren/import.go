package siren

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/jackc/pgx"
	"github.com/pkg/errors"

	"github.com/cheggaaa/pb"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	"fmt"
	"strings"
)

var cols = []string{
	"siret", "siren", "nic", "l1_normalisee", "l2_normalisee", "l3_normalisee", "l4_normalisee",
	"l5_normalisee", "l6_normalisee", "l7_normalisee", "l1_declaree", "l2_declaree", "l3_declaree", "l4_declaree",
	"l5_declaree", "l6_declaree", "l7_declaree", "numvoie", "indrep", "typvoie", "libvoie", "codpos", "cedex",
	"rpet", "libreg", "depet", "arronet", "ctonet", "comet", "libcom", "du", "tu", "uu", "epci", "tcd", "zemet",
	"siege", "enseigne", "ind_publipo", "diffcom", "amintret", "natetab", "libnatetab", "apet700", "libapet",
	"dapet", "tefet", "libtefet", "efetcent", "defet", "origine", "dcret", "ddebact", "activnat", "lieuact",
	"actisurf", "saisonat", "modet", "prodet", "prodpart", "auxilt", "nomen_long", "sigle", "nom", "prenom",
	"civilite", "rna", "nicsiege", "rpen", "depcomen", "adr_mail", "nj", "libnj", "apen700", "libapen", "dapen",
	"aprm", "ess", "dateess", "tefen", "libtefen", "efencent", "defen", "categorie", "dcren", "amintren", "monoact",
	"moden", "proden", "esaann", "tca", "esaapen", "esasec1n", "esasec2n", "esasec3n", "esasec4n", "vmaj", "vmaj1",
	"vmaj2", "vmaj3", "datemaj",
}

var majcols = append(cols, "eve", "dateve", "typcreh", "dreactet", "dreacten", "madresse", "menseigne",
	"mapet", "mprodet", "mauxilt", "mnomen", "msigle", "mnicsiege", "mnj", "mapen", "mproden", "siretps", "tel")

// CSVImport is a struct helper to import a CSV file to the database
// It implementes the pgx.Copy interface
type CSVImport struct {
	file   *os.File
	kind   FileType
	path   string
	bar    *pb.ProgressBar
	reader *csv.Reader
	values []interface{}
	err    error
}

// Copy actually copies the content of the CSV file to the database
func (c *CSVImport) Copy(db *pgx.ConnPool) error {
	switch c.kind {
	case StockType:
		c, err := db.CopyFrom(pgx.Identifier{"enterprises"}, cols, c)
		if err != nil {
			return errors.Wrap(err, "couldn't copyfrom")
		}
		logrus.WithField("records", c).Info("Imported Stock file")
	case DailyType:
		if _, err := db.Exec("TRUNCATE table temp_incremental"); err != nil {
			return errors.Wrap(err, "couldn't truncate table temp_incremental")
		}
		c, err := db.CopyFrom(pgx.Identifier{"temp_incremental"}, majcols, c)
		if err != nil {
			return errors.Wrap(err, "couldn't copyfrom")
		}
		if _, err = db.Exec("VACUUM ANALYZE temp_incremental"); err != nil {
			return errors.Wrap(err, "couldn't optimize table temp_incremental")
		}
		logrus.WithField("records", c).Info("Imported Daily file")
	}
	return nil
}

// Prepare opens the file and prepares the reader
func (c *CSVImport) Prepare() error {
	var err error
	var fd *os.File
	var fdi os.FileInfo

	if fd, err = os.Open(c.path); err != nil {
		return err
	}
	if fdi, err = fd.Stat(); err != nil {
		return err
	}
	c.file = fd
	c.reader = csv.NewReader(transform.NewReader(fd, charmap.Windows1252.NewDecoder()))
	c.reader.Comma = ';'

	c.bar = pb.New64(fdi.Size()).SetUnits(pb.U_BYTES)
	c.bar.ShowCounters = true
	c.bar.ShowPercent = true
	c.bar.ShowSpeed = true
	c.bar.ShowTimeLeft = true
	c.bar.Prefix("Importing " + c.path)
	c.bar.Start()
	c.reader.Read() // Skip the header part

	return nil
}

// Next returns true if there is another row and makes the next row data
// available to Values(). When there are no more rows available or an error
// has occurred it returns false.
// Satisifies the pgx.CopyFromSource interface
func (c *CSVImport) Next() bool {
	var err error
	var rec []string
	var siret string
	var values []interface{}

	if rec, err = c.reader.Read(); err != nil {
		if perr, ok := err.(*csv.ParseError); ok && perr.Err == csv.ErrFieldCount {
			return c.Next()
		}
		defer c.file.Close()
		defer c.bar.Finish()
		if err == io.EOF {
			return false
		}
		c.err = err
		return false
	}

	tot := 0
	siret = fmt.Sprintf("%s%s", rec[0], rec[1])
	values = append(values, siret)
	for _, v := range rec {
		values = append(values, v)
		tot += len(v) + 3 // two quotes and the ;
	}
	c.values = values
	c.bar.Add(tot - 3)

	return true
}

// Values returns the values for the current row.
func (c *CSVImport) Values() ([]interface{}, error) {
	return c.values, c.err
}

// Err returns any error that has been encountered by the CopyFromSource. If
// this is not nil *Conn.CopyFrom will abort the copy.
func (c *CSVImport) Err() error {
	return c.err
}

// Update update stock table from daily update file
func (c *CSVImport) Update(db *pgx.ConnPool) error {
	if c.kind == DailyType {
		var err error

		//Delete removed (E), exited (O) and modified (I) companies
		_, err = db.Exec(`
			DELETE FROM enterprises e
			USING temp_incremental i
			WHERE (e.siren, e.nic) = (i.siren, i.nic)
			AND i.vmaj IN ('E', 'O', 'I')
		`)
		if err != nil {
			return errors.Wrap(err, "couldn't remove enterprises from temp_incremental table")
		}

		//Adds new companies (C), which enter (D) and have been modified (F)
		var icols []string
		for _, col:= range cols {
			icols = append(icols, "i."+col)
		}
		_, err = db.Exec(fmt.Sprintf(
			"INSERT INTO enterprises (%s) SELECT %s FROM temp_incremental i WHERE i.vmaj in ('C', 'D', 'F')",
			strings.Join(cols, ", "),
			strings.Join(icols, ", "),
			))

		if err != nil {
			return errors.Wrap(err, "couldn't insert enterprises from temp_incremental table")
		}
	}

	return nil
}
