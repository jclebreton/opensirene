package logic

import (
	"github.com/jclebreton/opensirene/database"
	"github.com/jclebreton/opensirene/opendata/gouv_sirene"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// Import is the way to remote files to database
func Import(sfs gouv_sirene.RemoteFiles) error {
	var err error

	if err = database.InitImportClient(); err != nil {
		return errors.Wrap(err, "Couldn't initialize pgx")
	}

	//Lock database for import
	dbMutex := NewMutex(database.ImportClient)
	if err := dbMutex.Lock(); err != nil {
		return err
	}
	defer func() {
		err = dbMutex.Unlock()
		if err != nil {
			logrus.Warning(err)
		}
	}()

	//Download an extract
	if err = gouv_sirene.Do(sfs, 4); err != nil {
		return errors.Wrap(err, "Couldn't retrieve files")
	}

	cis, err := ToCSVImport(sfs)
	if err != nil {
		return errors.Wrap(err, "Couldn't convert to CSVImport")
	}

	tracker := NewTracker(database.ImportClient)
	if err = cis.Import(tracker); err != nil {
		return errors.Wrap(err, "Import error")
	}

	return nil
}
