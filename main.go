package main

import (
	"time"

	flag "github.com/ogier/pflag"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/jclebreton/opensirene/conf"
	"github.com/jclebreton/opensirene/database"
	"github.com/jclebreton/opensirene/download"
	"github.com/jclebreton/opensirene/opendata/siren"
	"github.com/jclebreton/opensirene/router"
)

func main() {
	var err error
	var cnf string
	var full bool

	flag.StringVarP(&cnf, "conf", "c", "conf.yml", "Path to the configuration file")
	flag.BoolVarP(&full, "full-import", "", false, "Get a full import from the last stock file")
	flag.Parse()

	if err = conf.Load(cnf); err != nil {
		logrus.WithError(err).Fatal("Couldn't parse configuration")
	}
	if full {
		if err = FullImport(); err != nil {
			logrus.WithError(err).Fatal("An error is occured during import")
		}
	}

	if err = database.InitQueryClient(); err != nil {
		logrus.WithError(err).Fatal("Couldn't initialize GORM")
	}
	defer database.DB.Close()

	if err = router.SetupAndRun(); err != nil {
		logrus.WithError(err).Fatal("Could not run the server")
	}
}

func FullImport() error {
	var err error
	var sfs siren.RemoteFiles

	s := time.Now()
	if err = database.InitImportClient(); err != nil {
		return errors.Wrap(err, "Couldn't initalize pgx")
	}

	if err = database.ImportClient.TryLock(); err != nil {
		return err
	}
	defer func() {
		err = database.ImportClient.Unlock()
		if err != nil {
			logrus.Warning(err)
		}
	}()

	if sfs, err = siren.GrabLatestFull(); err != nil {
		return errors.Wrap(err, "Couldn't grab full")
	}

	if err = download.Do(sfs, 4); err != nil {
		return errors.Wrap(err, "Couldn't retrieve files")
	}

	cis, err := sfs.ToCSVImport()
	if err != nil {
		return errors.Wrap(err, "Couldn't convert to CSVImport")
	}

	if err = cis.Import(); err != nil {
		return errors.Wrap(err, "Full import error")
	}

	logrus.WithField("took", time.Since(s)).Info("Done !")

	return nil
}
