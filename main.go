package main

import (
	"time"

	"github.com/Depado/lightsiren/conf"
	"github.com/Depado/lightsiren/opendata"
	flag "github.com/ogier/pflag"
	"github.com/sirupsen/logrus"
)

func main() {
	var err error
	var cnf string

	flag.StringVarP(&cnf, "conf", "c", "conf.yml", "Path to the configuration file")
	flag.Parse()

	if err = conf.Load(cnf); err != nil {
		logrus.WithError(err).Fatal("Couldn't parse configuration")
	}
	s := time.Now()
	if err = opendata.Grab(); err != nil {
		logrus.WithError(err).Fatal("Couldn't grab data")
	}
	logrus.WithField("took", time.Since(s)).Info("Done !")
}
