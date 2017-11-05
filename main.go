package main

import (
	"time"

	"github.com/jasonlvhit/gocron"
	"github.com/jinzhu/gorm"

	flag "github.com/ogier/pflag"
	"github.com/sirupsen/logrus"

	"github.com/jclebreton/opensirene/api/router"
	"github.com/jclebreton/opensirene/conf"
	"github.com/jclebreton/opensirene/database"
	"github.com/jclebreton/opensirene/logic"
	"github.com/jclebreton/opensirene/opendata/gouvfr/sirene"
)

func main() {
	var err error
	var config string
	var fullImport bool
	var pgxClient *database.PgxClient
	var gormClient *gorm.DB

	// Configuration
	flag.StringVarP(&config, "conf", "c", "conf.yml", "Path to the configuration file")
	flag.BoolVarP(&fullImport, "full-import", "", false, "Get a full import from the last stock file")
	flag.Parse()
	if err = conf.Load(config); err != nil {
		logrus.WithError(err).Fatal("Couldn't parse configuration")
	}

	// Init PGX database client
	if pgxClient, err = database.NewImportClient(); err != nil {
		logrus.WithError(err).Fatal("Couldn't initialize PGX client")
	}

	// Full import
	if fullImport {
		s := time.Now()
		var sfs sirene.RemoteFiles
		if sfs, err = sirene.GrabLatestFull(); err != nil {
			logrus.WithError(err).Fatal("An error is occured during grab")
		}

		if err = logic.Import(pgxClient, sfs); err != nil {
			logrus.WithError(err).Fatal("An error is occurred during full import")
		}
		logrus.WithField("import took", time.Since(s)).Info("Done !")
	}

	// Init GORM database client
	if gormClient, err = database.NewGORMClient(); err != nil {
		logrus.WithError(err).Fatal("Couldn't initialize GORM")
	}
	defer gormClient.Close()

	//Enable update as crontab
	go func() {
		gocron.Every(3).Hours().Do(logic.Daily, pgxClient)
		// Execute the update at startup
		gocron.RunAll()
		_, t := gocron.NextRun()
		logrus.WithField("next", t).Info("Started cron background task")
		<-gocron.Start()
	}()

	//Start API
	if err = router.SetupAndRun(gormClient); err != nil {
		logrus.WithError(err).Fatal("Could not run the server")
	}
}
