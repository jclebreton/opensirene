package main

import (
	"github.com/Depado/lightsiren/conf"
	"github.com/Depado/lightsiren/download"
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
	download.Start()
}
