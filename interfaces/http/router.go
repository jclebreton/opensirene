package http

//TODO : remove all references to conf
import (
	"fmt"

	"github.com/Depado/ginprom"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/jclebreton/opensirene/conf"
)

type ServerConfigurer interface {
	CorsEnabled() bool
	CorsPermissive() bool
	DebugMode() bool
}

func NewServer(sc ServerConfigurer) server {
	return server{
		cors:           sc.CorsEnabled(),
		corsPermissive: sc.CorsPermissive(),
		debug:          sc.DebugMode(),
	}
}

type server struct {
	cors           bool
	corsPermissive bool
	debug          bool
	ginEngine      *gin.Engine
}

func (s *server) SetupRouter() {
	// Create the router
	s.ginEngine = gin.Default()

	// Setup debug mode or not in Gin
	if s.debug {
		gin.SetMode(gin.ReleaseMode)

	}

	// Setup CORS configuration
	if s.cors {
		cc := cors.Config{
			AllowMethods:  []string{"POST", "GET", "OPTIONS", "PUT", "DELETE", "PATCH"},
			AllowHeaders:  []string{"DNT", "X-CustomHeader", "Keep-Alive", "User-Agent", "X-Requested-With", "If-Modified-Since", "Cache-Control", "Content-Type", "Content-Range", "Range", "Authorization"},
			ExposeHeaders: []string{"DNT", "X-CustomHeader", "Keep-Alive", "User-Agent", "X-Requested-With", "If-Modified-Since", "Cache-Control", "Content-Type", "Content-Range", "Range", "Authorization"},
		}
		if s.corsPermissive {
			cc.AllowAllOrigins = true
		} else {
			cc.AllowOrigins = conf.C.Server.Cors.AllowOrigins
		}
		s.ginEngine.Use(cors.New(cc))
	}

	p := ginprom.New(ginprom.Subsystem(conf.C.Prometheus.Prefix), ginprom.Engine(s.ginEngine))

	s.ginEngine.Use(p.Instrument())
}

func (s server) SetupRoutes(httpG HttpGateway) {
	s.ginEngine.GET("/history", httpG.GetHistories)
	//r.GET("/siret/:id", views.GetSiret)
	//r.GET("/siren/:id", views.GetSiren)
}

func (s server) Start() error {
	// Run the server
	logrus.WithFields(logrus.Fields{"port": conf.C.Server.Port, "host": conf.C.Server.Host}).Info("Starting server")
	return s.ginEngine.Run(fmt.Sprintf("%s:%d", conf.C.Server.Host, conf.C.Server.Port))
}
