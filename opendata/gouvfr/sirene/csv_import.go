package sirene

import (
	"fmt"
	"os"
	"strings"

	"github.com/jackc/pgx"
	"github.com/jclebreton/opensirene/database"
	"github.com/jclebreton/opensirene/progress"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var cols = []string{
	"siren", "nic", "l1_normalisee", "l2_normalisee", "l3_normalisee", "l4_normalisee",
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
	Path         string
	File         *os.File
	Kind         FileType
	ZipName      string
	ProgressChan chan *progress.Progress
}

// Copy actually copies the content of the CSV file to the database
func (c *CSVImport) Copy(db *pgx.ConnPool) error {
	var err error
	var total int

	// Prepare SQL Copy
	cf := &database.PgxCopyFrom{
		Path: c.Path,
		File: c.File,
		CallBackTriggerOnColName: []string{"DAPET", "DEFET", "DAPEN", "DEFEN", "ESAANN", "DATEMAJ", "AMINTRET", "AMINTREN",
			"DDEBACT", "DATEESS", "DATEVE", "DREACTET", "DREACTEN", "DCRET", "DCREN"},
		CallBackFunc: func(colValue string) (interface{}, error) {
			if colValue != "" {
				d, err := NewDateSirene(colValue)
				if err != nil {
					return nil, errors.Wrap(err, "couldn't convert date")
				}
				return d.GetDate(), nil
			}
			return nil, nil
		},
	}
	if err = cf.Prepare(); err != nil {
		return errors.Wrap(err, "Couldn't prepare import")
	}

	switch c.Kind {
	case StockType:
		total, err = c.copyFile(db, "enterprises", cols, cf, false)
	case DailyType:
		total, err = c.copyFile(db, "temp_incremental", majcols, cf, true)
	default:
		return errors.New("couldn't import unknown type file!")
	}

	if err != nil {
		return errors.Wrap(err, "couldn't copy file to database")
	}

	logrus.WithFields(logrus.Fields{"records": c, "total": total}).Info("Imported Stock file")

	return nil
}

// Copy actually copies the content of the CSV file to the database
// Empty table before processing and optimize it after
func (c *CSVImport) copyFile(db *pgx.ConnPool, tableName string, cols []string, cf *database.PgxCopyFrom, optimize bool) (int, error) {
	// Reset table
	if _, err := db.Exec(fmt.Sprintf("TRUNCATE table %s", tableName)); err != nil {
		return 0, errors.Wrap(err, fmt.Sprintf("couldn't truncate %s", tableName))
	}

	//Copy
	total, err := db.CopyFrom(pgx.Identifier{tableName}, cols, cf)
	if err != nil {
		return total, errors.Wrap(err, "couldn't import csv to database")
	}

	//Optimize table performance
	if optimize {
		if _, err = db.Exec(fmt.Sprintf("VACUUM ANALYZE %s", tableName)); err != nil {
			return total, errors.Wrap(err, fmt.Sprintf("couldn't optimize %s", tableName))
		}
	}

	return total, nil
}

// Update update stock table from daily update file
func (c *CSVImport) Update(db *pgx.ConnPool) error {
	if c.Kind == DailyType {
		var err error

		// Create transaction to avoid corruptions
		transaction, err := db.Begin()
		if err != nil {
			return errors.Wrap(err, "couldn't begin sql transaction")
		}

		defer transaction.Rollback()

		// Delete removed (E), exited (O) and modified (I) companies
		_, err = transaction.Exec(`
			DELETE FROM enterprises e
			USING temp_incremental i
			WHERE (e.siren, e.nic) = (i.siren, i.nic)
			AND i.vmaj IN ('E', 'O', 'I')
		`)
		if err != nil {
			return errors.Wrap(err, "couldn't remove enterprises from temp_incremental table")
		}

		// Adds new companies (C), which enter (D) and have been modified (F)
		var icols []string
		for _, col := range cols {
			icols = append(icols, "i."+col)
		}
		_, err = transaction.Exec(fmt.Sprintf(
			"INSERT INTO enterprises (%s) SELECT %s FROM temp_incremental i WHERE i.vmaj in ('C', 'D', 'F')",
			strings.Join(cols, ", "),
			strings.Join(icols, ", "),
		))
		if err != nil {
			return errors.Wrap(err, "couldn't insert enterprises from temp_incremental table")
		}

		transaction.Commit()
	}

	return nil
}
