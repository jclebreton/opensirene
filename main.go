package main

import (
	"time"

	flag "github.com/ogier/pflag"
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
	var sfs siren.RemoteFiles

	flag.StringVarP(&cnf, "conf", "c", "conf.yml", "Path to the configuration file")
	flag.BoolVarP(&full, "full-import", "", false, "Get a full import from the last stock file")
	flag.Parse()

	if err = conf.Load(cnf); err != nil {
		logrus.WithError(err).Fatal("Couldn't parse configuration")
	}
	if full {
		s := time.Now()
		if err = database.InitImportClient(); err != nil {
			logrus.WithError(err).Fatal("Couldn't initalize pgx")
		}
		if sfs, err = siren.GrabLatestFull(); err != nil {
			logrus.WithError(err).Fatal("Couldn't grab full")
		}
		if err = download.Do(sfs, 4); err != nil {
			logrus.WithError(err).Fatal("Couldn't retrieve files")
		}
		cis, err := sfs.ToCSVImport()
		if err != nil {
			logrus.WithError(err).Fatal("Couldn't convert to CSVImport")
		}
		for _, ci := range cis {
			if err = ci.Prepare(); err != nil {
				logrus.WithError(err).Fatal("Couldn't prepare import")
			}
			if err = ci.Copy(database.ImportClient.Conn); err != nil {
				logrus.WithError(err).Fatal("Couldn't copy")
			}
			if err = ci.Update(database.ImportClient.Conn); err != nil {
				logrus.WithError(err).Fatal("Couldn't apply update")
			}
		}
		logrus.WithField("took", time.Since(s)).Info("Done !")
	}

	if err = database.InitQueryClient(); err != nil {
		logrus.WithError(err).Fatal("Couldn't initialize GORM")
	}
	defer database.DB.Close()

	if err = router.SetupAndRun(); err != nil {
		logrus.WithError(err).Fatal("Could not run the server")
	}
}
