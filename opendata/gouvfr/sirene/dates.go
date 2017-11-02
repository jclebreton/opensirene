package sirene

import (
	"errors"
	"time"
)

var layouts = [4]string{"2006-01-02T15:04:05", "20060102", "200601", "2006"}
var toStringFormat = "01-02-2006" // PostgreSQL default DateStyle: ISO, MDY

type dateSirene struct {
	raw       string
	converted time.Time
}

func NewDateSirene(raw string) (dateSirene, error) {
	ds := &dateSirene{raw: raw}
	err := ds.convert()
	if err != nil {
		return dateSirene{}, err
	}

	return *ds, nil
}

func (ds *dateSirene) convert() error {
	for _, layout := range layouts {
		if d, err := time.Parse(layout, ds.raw); err == nil {
			ds.converted = d
			return nil
		}
	}
	return errors.New("couldn't convert date")
}

// GetDate returns the converted date
func (ds *dateSirene) GetDate() time.Time {
	return ds.converted
}

// String returns the current date as string
func (c *dateSirene) String() string {
	return c.converted.Format(toStringFormat)
}
