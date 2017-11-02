package logic

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/jclebreton/opensirene/database"
	"github.com/jclebreton/opensirene/opendata/gouvfr/sirene"
)

// Import is the way to remote files to database
func Import(sfs sirene.RemoteFiles) error {
	var err error

	if err = database.InitImportClient(); err != nil {
		return errors.Wrap(err, "Couldn't initialize pgx")
	}

	//Lock database for import
	dbMutex := newMutex(database.ImportClient)
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
	if err = sirene.Do(sfs, 4); err != nil {
		return errors.Wrap(err, "Couldn't retrieve files")
	}

	cis, err := ToCSVImport(sfs)
	if err != nil {
		return errors.Wrap(err, "Couldn't convert to CSVImport")
	}

	tracker := newTracker(database.ImportClient)
	if err = cis.Import(tracker); err != nil {
		return errors.Wrap(err, "Import error")
	}

	return nil
}
