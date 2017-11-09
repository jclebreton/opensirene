package main

import (
	"github.com/jinzhu/gorm"
	flag "github.com/ogier/pflag"

	"github.com/jackc/pgx"
	"github.com/jclebreton/opensirene/conf"
	"github.com/jclebreton/opensirene/database"
	http "github.com/jclebreton/opensirene/interfaces/http"
	cors "github.com/jclebreton/opensirene/interfaces/http/cors"
	"github.com/jclebreton/opensirene/interfaces/http/monitoring"
	"github.com/jclebreton/opensirene/interfaces/json"
	"github.com/jclebreton/opensirene/interfaces/storage/db_status"
	"github.com/jclebreton/opensirene/interfaces/storage/establishments"
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

	pgxClient, err := database.NewPgxClientClient()
	if err != nil {
		logrus.WithError(err).Fatal("Couldn't initialize Pgx")
	}
	defer pgxClient.Close()

	server := setupServer(gormClient, pgxClient)
	if err := server.Start(conf.C.Server); err != nil {
		logrus.WithError(err).Fatal("Could not setup and run API")
	}
}

func loadConf() {
	var err error
	var config string
	var fullImport bool

	flag.StringVarP(&config, "config", "c", "conf.yml", "Path to the configuration file")
	flag.BoolVarP(&fullImport, "drop", "", false, "Truncate database and run a full import")
	flag.Parse()
	if err = conf.Load(config); err != nil {
		logrus.WithError(err).Fatal("Couldn't parse configuration")
	}
}

func setupServer(gormClient *gorm.DB, pgxClient *pgx.ConnPool) http.Server {
	server := http.NewServer(conf.C.Server)
	server.SetupRouter()
	server.StartMonitoring(monitoring.NewPrometheus(conf.C.Prometheus.Prefix, server.GinEngine))
	server.SetupRoutes(http.NewHttpGateway(setInteractor(gormClient, pgxClient)), conf.C.Server.Prefix)
	if conf.C.Server.Cors.Enabled {
		server.SetupCors(cors.NewStandardCors(conf.C.Server.Cors.PermissiveMode, conf.C.Server.Cors.AllowOrigins))
	}
	return server
}

// set here the structs you want to implement the interfaces
func setInteractor(gormClient *gorm.DB, pgxClient *pgx.ConnPool) usecases.Interactor {
	return usecases.NewInteractor(
		&db_status.RW{PgxClient: pgxClient},
		&establishments.RW{GormClient: gormClient, PgxClient: pgxClient},
		&json.JSONwriterStd{},
	)
}
