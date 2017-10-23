package main

import (
	"time"

	flag "github.com/ogier/pflag"
	"github.com/sirupsen/logrus"

	"github.com/Depado/lightsiren/conf"
	"github.com/Depado/lightsiren/download"
	"github.com/Depado/lightsiren/opendata/siren"
)

func main() {
	var err error
	var cnf string
	var sfs siren.RemoteFiles

	flag.StringVarP(&cnf, "conf", "c", "conf.yml", "Path to the configuration file")
	flag.Parse()

	if err = conf.Load(cnf); err != nil {
		logrus.WithError(err).Fatal("Couldn't parse configuration")
	}
	s := time.Now()
	if sfs, err = siren.GrabLatestFull(); err != nil {
		logrus.WithError(err).Fatal("Couldn't grab full")
	}
	if err = download.Do(sfs, 4); err != nil {
		logrus.WithError(err).Fatal("Couldn't retrieve files")
	}
	logrus.WithField("took", time.Since(s)).Info("Done !")
}
