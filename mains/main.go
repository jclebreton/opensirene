package main

import (
	"github.com/jinzhu/gorm"
	flag "github.com/ogier/pflag"

	"github.com/jclebreton/opensirene/conf"
	"github.com/jclebreton/opensirene/database"
	http "github.com/jclebreton/opensirene/interfaces/http"
	"github.com/jclebreton/opensirene/interfaces/json"
	"github.com/jclebreton/opensirene/interfaces/storage/history"
	"github.com/jclebreton/opensirene/usecases"
	"github.com/sirupsen/logrus"
)

func main() {
	loadConf()

	gormClient, err := database.NewGORMClientFromString(conf.C.Database.ConnectionString())
	if err != nil {
		logrus.WithError(err).Fatal("Couldn't initialize GORM")
	}
	defer func() {
		err := gormClient.Close()
		logrus.WithError(err).Fatal("Couldn't close GORM")
	}()

	server := http.NewServer(conf.C.Server)
	server.SetupRouter()
	server.SetupRoutes(http.NewHttpGateway(setInteractor(gormClient)))

	if err := server.Start(); err != nil {
		logrus.WithError(err).Fatal("Could not setup and run API")
	}
}

// set here the structs you want to implement the interfaces
func setInteractor(db *gorm.DB) usecases.Interactor {
	return usecases.NewInteractor(
		history.RW{GormClient: db},
		json.JSONwriterStd{},
	)
}

func loadConf() {
	var err error
	var config string
	var fullImport bool

	// Configuration
	flag.StringVarP(&config, "config", "c", "conf.yml", "Path to the configuration file")
	flag.BoolVarP(&fullImport, "drop", "", false, "Truncate database and run a full import")
	flag.Parse()
	if err = conf.Load(config); err != nil {
		logrus.WithError(err).Fatal("Couldn't parse configuration")
	}
}
