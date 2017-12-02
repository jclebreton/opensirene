package gin

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jclebreton/opensirene/conf"
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
	GetApiPrefix() string
	GetAdminPrefix() string
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

func (s Server) SetupRoutes(httpG HttpGateway, prefix conf.Prefix) {
	s.GinEngine.GET(prefix.Admin+"/ping", httpG.Ping)
	s.GinEngine.GET(prefix.Admin+"/health", httpG.Health)
	s.GinEngine.GET(prefix.Admin+"/history", httpG.GetDBStatus)

	s.GinEngine.GET(prefix.Api+"/siret/:id", httpG.GetEnterpriseFromSiret)
	s.GinEngine.GET(prefix.Api+"/siren/:id", httpG.GetEstablishmentsFromSiren)
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
