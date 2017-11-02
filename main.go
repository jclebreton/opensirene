package main

import (
	"time"

	"github.com/jasonlvhit/gocron"

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
		var sfs sirene.RemoteFiles
		if sfs, err = sirene.GrabLatestFull(); err != nil {
			logrus.WithError(err).Fatal("An error is occured during grab")
		}

		if err = logic.Import(sfs); err != nil {
			logrus.WithError(err).Fatal("An error is occurred during full import")
		}
		logrus.WithField("import took", time.Since(s)).Info("Done !")
	}

	if err = database.InitQueryClient(); err != nil {
		logrus.WithError(err).Fatal("Couldn't initialize GORM")
	}
	defer database.DB.Close()

	go func() {
		gocron.Every(3).Hours().Do(logic.Daily)
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
