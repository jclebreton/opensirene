package main

import (
	"time"

	"github.com/jasonlvhit/gocron"

	flag "github.com/ogier/pflag"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/jclebreton/opensirene/api/router"
	"github.com/jclebreton/opensirene/conf"
	"github.com/jclebreton/opensirene/database"
	"github.com/jclebreton/opensirene/logic"
	"github.com/jclebreton/opensirene/opendata/gouv_sirene"
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
		s := time.Now()
		var sfs gouv_sirene.RemoteFiles
		if sfs, err = gouv_sirene.GrabLatestFull(); err != nil {
			logrus.WithError(err).Fatal("An error is occured during grab")
		}

		if err = Import(sfs); err != nil {
			logrus.WithError(err).Fatal("An error is occurred during full import")
		}
		logrus.WithField("import took", time.Since(s)).Info("Done !")
	}

	if err = database.InitQueryClient(); err != nil {
		logrus.WithError(err).Fatal("Couldn't initialize GORM")
	}
	defer database.DB.Close()

	go func() {
		gocron.Every(3).Hours().Do(Daily)
		// Execute the update at startup
		gocron.RunAll()
		_, t := gocron.NextRun()
		logrus.WithField("next", t).Info("Started cron background task")
		<-gocron.Start()
	}()

	if err = router.SetupAndRun(); err != nil {
		logrus.WithError(err).Fatal("Could not run the server")
	}
}

// Daily is the cron task that runs every few hours to get and apply the latest
// updates
func Daily() {
	var err error
	var sfs gouv_sirene.RemoteFiles

	if sfs, err = gouv_sirene.GrabLatestFull(); err != nil {
		logrus.WithError(err).Error("Could not download latest")
		return
	}

	sfs = sfs.Diff(logic.GetSuccessfulUpdateList())

	if err = Import(sfs); err != nil {
		logrus.WithError(err).Error("Could not download latest")
		return
	}
}

// Import is the way to remote files to database
func Import(sfs gouv_sirene.RemoteFiles) error {
	var err error

	if err = database.InitImportClient(); err != nil {
		return errors.Wrap(err, "Couldn't initalize pgx")
	}

	dbMutex := logic.NewMutex(database.ImportClient)
	if err := dbMutex.Lock(); err != nil {
		return err
	}
	defer func() {
		err = dbMutex.Unlock()
		if err != nil {
			logrus.Warning(err)
		}
	}()

	if err = logic.Do(sfs, 4); err != nil {
		return errors.Wrap(err, "Couldn't retrieve files")
	}

	cis, err := sfs.ToCSVImport()
	if err != nil {
		return errors.Wrap(err, "Couldn't convert to CSVImport")
	}

	if err = cis.Import(); err != nil {
		return errors.Wrap(err, "Import error")
	}

	return nil
}
