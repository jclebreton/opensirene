package main

import (
	"fmt"
	"time"

	flag "github.com/ogier/pflag"
	"github.com/sirupsen/logrus"

	"github.com/Depado/lightsiren/conf"
	"github.com/Depado/lightsiren/opendata"
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
	var d *opendata.Dataset
	if d, err = opendata.Grab(); err != nil {
		logrus.WithError(err).Fatal("Couldn't grab data")
	}
	for _, v := range d.Resources {
		fmt.Println(v)
	}
	logrus.WithField("took", time.Since(s)).Info("Done !")
}
