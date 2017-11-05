package main

import (
	"time"

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

	// Configuration
	flag.StringVarP(&config, "conf", "c", "conf.yml", "Path to the configuration file")
	flag.BoolVarP(&fullImport, "full-import", "", false, "Get a full import from the last stock file")
	flag.Parse()
	if err = conf.Load(config); err != nil {
		logrus.WithError(err).Fatal("Couldn't parse configuration")
	}

	// Init PGX database client
	var pgxClient *database.PgxClient
	if pgxClient, err = database.NewImportClient(); err != nil {
		logrus.WithError(err).Fatal("Couldn't initialize PGX client")
	}
	defer pgxClient.Conn.Close()

	// Full import
	if fullImport {
		s := time.Now()
		var sfs sirene.RemoteFiles
		if sfs, err = sirene.GrabLatestFull(); err != nil {
			logrus.WithError(err).Fatal("An error is occured during grab")
		}

		if err = logic.ImportRemoteFiles(pgxClient, sfs); err != nil {
			logrus.WithError(err).Fatal("An error is occurred during full import")
		}
		logrus.WithField("import took", time.Since(s)).Info("Done !")
	}

	//Start automatic updates
	crontab := &logic.Crontab{PgxClient: pgxClient}
	go crontab.Start()

	//Start API
	var gormClient *gorm.DB
	if gormClient, err = database.NewGORMClient(); err != nil {
		logrus.WithError(err).Fatal("Couldn't initialize GORM")
	}
	defer gormClient.Close()
	if err = router.SetupAndRun(gormClient); err != nil {
		logrus.WithError(err).Fatal("Could not setup and run API")
	}
}
