package logic

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/jclebreton/opensirene/database"
	"github.com/jclebreton/opensirene/opendata/gouvfr/sirene"
)

// Import is the way to remote files to database
func ImportRemoteFiles(pgxClient *database.PgxClient, remoteFiles sirene.RemoteFiles) error {
	var err error

	//Lock database for import
	dbMutex := database.NewPgxMutex(pgxClient)
	if err = dbMutex.Lock(); err != nil {
		return err
	}
	defer func() {
		if err = dbMutex.Unlock(); err != nil {
			logrus.WithError(err).Warning("Couldn't freeze database mutex")
		}
	}()

	//Download an extract
	if err = sirene.Do(remoteFiles, 4); err != nil {
		return errors.Wrap(err, "Couldn't retrieve files")
	}

	// Convert them
	cis, err := toCSVImport(remoteFiles)
	if err != nil {
		return errors.Wrap(err, "Couldn't convert to CSVImport")
	}

	//Import
	tracker := NewTracker(pgxClient)
	if err = cis.importCSVFiles(pgxClient, tracker); err != nil {
		return errors.Wrap(err, "Import error")
	}

	return nil
}

// ToCSVImport converts a slice of RemoteFile to a slice of CSVImport.
// It expects that at least one file was extracted
func toCSVImport(rfs sirene.RemoteFiles) (CSVImports, error) {
	var out CSVImports
	for _, rf := range rfs {
		out = append(out, &sirene.CSVImport{
			Path:    rf.ExtractedFiles[0],
			Kind:    rf.Type,
			ZipName: rf.FileName,
		})
	}
	return out, nil
}

// CSVImports is a slice of pointer to CSVImport
type CSVImports []*sirene.CSVImport

// Import will import each CSVImport present in the slice
func (c CSVImports) importCSVFiles(pgxClient *database.PgxClient, tracker Tracker) error {
	var err error
	for _, ci := range c {

		if ci.Kind == sirene.StockType {
			if err = tracker.Truncate(); err != nil {
				return errors.Wrap(err, "Unable to reset history before import stock file")
			}
		}

		if err = ci.Copy(pgxClient.Conn); err != nil {
			if e := tracker.Save(ci.ZipName, false, err.Error()); e != nil {
				return errors.Wrap(err, e.Error())
			}
			return errors.Wrap(err, "Couldn't copy")
		}

		if err = ci.Update(pgxClient.Conn); err != nil {
			if e := tracker.Save(ci.ZipName, false, err.Error()); e != nil {
				return errors.Wrap(err, e.Error())
			}
			return errors.Wrap(err, "Couldn't apply update")
		}

		// Save step to db
		err = tracker.Save(ci.ZipName, true, "")
		if err != nil {
			return errors.Wrap(err, "Couldn't log")
		}
	}
	return nil
}
