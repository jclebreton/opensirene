package http

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Server struct {
	debug     bool
	GinEngine *gin.Engine
}

type ServerConfigurer interface {
	GetPort() int
	GetHost() string
	DebugMode() bool
}

func NewServer(sc ServerConfigurer) Server {
	return Server{
		debug: sc.DebugMode(),
	}
}

func (s Server) Start(conf ServerConfigurer) error {
	logrus.WithFields(logrus.Fields{"port": conf.GetPort(), "host": conf.GetHost()}).Info("Starting server")
	return s.GinEngine.Run(fmt.Sprintf("%s:%d", conf.GetHost(), conf.GetPort()))
}

func (s *Server) SetupRouter() {
	s.GinEngine = gin.Default()
	if s.debug {
		gin.SetMode(gin.ReleaseMode)
	}
}

func (s Server) SetupRoutes(httpG HttpGateway) {
	s.GinEngine.GET("/history", httpG.GetHistories)
	//r.GET("/siret/:id", views.GetSiret)
	//r.GET("/siren/:id", views.GetSiren)
}

type Monitorer interface {
	Monitor() gin.HandlerFunc
}

func (s *Server) StartMonitoring(monitorer Monitorer) {
	s.GinEngine.Use(monitorer.Monitor())
}

type CorsSetter interface {
	SetCors() gin.HandlerFunc
}

func (s *Server) SetupCors(setter CorsSetter) {
	s.GinEngine.Use(setter.SetCors())
}
