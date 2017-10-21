package main

import (
	"time"

	flag "github.com/ogier/pflag"
	"github.com/sirupsen/logrus"

	"github.com/Depado/lightsiren/conf"
	"github.com/Depado/lightsiren/download"
	"github.com/Depado/lightsiren/opendata"
)

func main() {
	var err error
	var cnf string
	var sfs opendata.SireneFiles

	flag.StringVarP(&cnf, "conf", "c", "conf.yml", "Path to the configuration file")
	flag.Parse()

	if err = conf.Load(cnf); err != nil {
		logrus.WithError(err).Fatal("Couldn't parse configuration")
	}
	s := time.Now()
	if sfs, err = opendata.GrabLatestFull(); err != nil {
		logrus.WithError(err).Fatal("Couldn't grab full")
	}
	download.Do(sfs)
	logrus.WithField("took", time.Since(s)).Info("Done !")
}
