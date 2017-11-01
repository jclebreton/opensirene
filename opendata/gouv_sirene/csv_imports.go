package gouv_sirene

import (
	"github.com/jclebreton/opensirene/crontab"
	"github.com/jclebreton/opensirene/database"
	"github.com/pkg/errors"
)

// CSVImports is a slice of pointer to CSVImport
type CSVImports []*CSVImport

// Import will import each CSVImport present in the slice
func (c CSVImports) Import(tracker crontab.Tracker) error {
	var err error
	for _, ci := range c {
		if err = ci.Prepare(); err != nil {
			if e := tracker.Save(FileTypeName(ci.Kind), err.Error(), ci.ZipName, false); e != nil {
				return errors.Wrap(err, e.Error())
			}
			return errors.Wrap(err, "Couldn't prepare import")
		}
		if err = ci.Copy(database.ImportClient.Conn); err != nil {
			if e := tracker.Save(FileTypeName(ci.Kind), err.Error(), ci.ZipName, false); e != nil {
				return errors.Wrap(err, e.Error())
			}
			return errors.Wrap(err, "Couldn't copy")
		}
		if err = ci.Update(database.ImportClient.Conn); err != nil {
			if e := tracker.Save(FileTypeName(ci.Kind), err.Error(), ci.ZipName, false); e != nil {
				return errors.Wrap(err, e.Error())
			}
			return errors.Wrap(err, "Couldn't apply update")
		}

		err = tracker.Save(FileTypeName(ci.Kind), "", ci.ZipName, true)
		if err != nil {
			return errors.Wrap(err, "Couldn't log")
		}
	}
	return nil
}
