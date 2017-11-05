package logic

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/jclebreton/opensirene/database"
	"github.com/jclebreton/opensirene/opendata/gouvfr/sirene"
)

// Import is the way to remote files to database
func Import(pgxClient *database.PgxClient, remoteFiles sirene.RemoteFiles) error {
	var err error

	//Lock database for import
	dbMutex := newMutex(pgxClient)
	if err = dbMutex.Lock(); err != nil {
		return err
	}
	defer func() {
		perr := dbMutex.Unlock()
		if perr != nil {
			logrus.Warning(perr)
		}
	}()

	//Download an extract
	if err = sirene.Do(remoteFiles, 4); err != nil {
		return errors.Wrap(err, "Couldn't retrieve files")
	}

	// Convert them
	cis, err := ToCSVImport(remoteFiles)
	if err != nil {
		return errors.Wrap(err, "Couldn't convert to CSVImport")
	}

	//Import
	tracker := newTracker(pgxClient)
	if err = cis.Import(pgxClient, tracker); err != nil {
		return errors.Wrap(err, "Import error")
	}

	return nil
}
