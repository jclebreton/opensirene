package database

import (
	"encoding/csv"
	"io"
	"os"

	"time"

	"github.com/cheggaaa/pb"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

// pgxCopyFrom is a struct helper to import a CSV file to the database
// It implement's the pgx.Copy interface
type PgxCopyFrom struct {
	values []interface{}
	err    error
	reader *csv.Reader

	Path string
	File *os.File
	Bar  *pb.ProgressBar

	headers               []string
	callBackTriggerOnKeys []int

	CallBackTriggerOnColName []string
	CallBackFunc             func(colValue string) (interface{}, error)
}

// Prepare opens the file and prepares the reader
func (c *PgxCopyFrom) Prepare() error {
	var err error
	var fd *os.File
	var fdi os.FileInfo

	if fd, err = os.Open(c.Path); err != nil {
		return err
	}
	if fdi, err = fd.Stat(); err != nil {
		return err
	}
	c.File = fd
	c.reader = csv.NewReader(transform.NewReader(fd, charmap.Windows1252.NewDecoder()))
	c.reader.Comma = ';'

	c.Bar = pb.New64(fdi.Size()).SetUnits(pb.U_BYTES)
	c.Bar.ShowCounters = true
	c.Bar.ShowPercent = true
	c.Bar.ShowSpeed = true
	c.Bar.ShowTimeLeft = true
	c.Bar.Prefix("Importing " + c.Path)
	c.Bar.Start()

	// Save and skip the header part
	c.headers, err = c.reader.Read()
	if err != nil {
		return err
	}

	//Search columns indexes for callback
	for _, colTriggerName := range c.CallBackTriggerOnColName {
		for k, v := range c.headers {
			if colTriggerName == v {
				c.callBackTriggerOnKeys = append(c.callBackTriggerOnKeys, k)
				break
			}
		}
	}

	return nil
}

func (c *PgxCopyFrom) colHasTrigger(k int) bool {
	for _, v := range c.callBackTriggerOnKeys {
		if k == v {
			return true
		}
	}
	return false
}

func (c *PgxCopyFrom) callTrigger(values []interface{}, k int, v string) ([]interface{}, error) {
	if c.colHasTrigger(k) {
		mixed, err := c.CallBackFunc(v)
		if err != nil {
			return nil, err
		}
		if t, ok := mixed.(time.Time); ok {
			values = append(values, t)
		} else {
			values = append(values, mixed)
		}
	} else {
		values = append(values, v)
	}
	return values, nil
}

// Next returns true if there is another row and makes the next row data
// available to Values(). When there are no more rows available or an error
// has occurred it returns false.
// Satisfies the pgx.CopyFromSource interface
func (c *PgxCopyFrom) Next() bool {
	var err error
	var rec []string
	var values []interface{}

	if rec, err = c.reader.Read(); err != nil {
		if perr, ok := err.(*csv.ParseError); ok && perr.Err == csv.ErrFieldCount {
			return c.Next()
		}
		defer c.File.Close()
		defer c.Bar.Finish()
		if err == io.EOF {
			return false
		}
		c.err = err
		return false
	}

	tot := 0
	for k, v := range rec {
		tot += len(v) + 3 // two quotes and the ;
		if values, err = c.callTrigger(values, k, v); err != nil {
			c.err = err
			return false
		}
	}

	c.values = values
	c.Bar.Add(tot - 3) // the last ; and the \n

	return true
}

// Values returns the values for the current row.
// Satisfies the pgx.CopyFromSource interface
func (c *PgxCopyFrom) Values() ([]interface{}, error) {
	return c.values, c.err
}

// Err returns any error that has been encountered by the CopyFromSource. If
// this is not nil *Conn.CopyFrom will abort the copy.
// Satisfies the pgx.CopyFromSource interface
func (c *PgxCopyFrom) Err() error {
	return c.err
}
