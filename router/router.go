package router

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/jclebreton/opensirene/conf"
	"github.com/jclebreton/opensirene/views"
)

// SetupAndRun creates the router and runs it
func SetupAndRun() error {
	var err error

	// Create the router
	r := gin.Default()

	// Setup debug mode or not in Gin
	if !conf.C.Server.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	// Setup CORS configuration
	cc := cors.Config{
		AllowMethods:  []string{"POST", "GET", "OPTIONS", "PUT", "DELETE", "PATCH"},
		AllowHeaders:  []string{"DNT", "X-CustomHeader", "Keep-Alive", "User-Agent", "X-Requested-With", "If-Modified-Since", "Cache-Control", "Content-Type", "Content-Range", "Range", "Authorization"},
		ExposeHeaders: []string{"DNT", "X-CustomHeader", "Keep-Alive", "User-Agent", "X-Requested-With", "If-Modified-Since", "Cache-Control", "Content-Type", "Content-Range", "Range", "Authorization"},
	}
	if conf.C.Server.Cors.PermissiveMode {
		cc.AllowAllOrigins = true
	} else {
		cc.AllowOrigins = conf.C.Server.Cors.AllowOrigins
	}
	r.Use(cors.New(cc))

	// Route setup
	r.GET("/siret/:id", views.GetSiret)
	r.GET("/siren/:id", views.GetSiren)

	// Run the server
	logrus.WithFields(logrus.Fields{"port": conf.C.Server.Port, "host": conf.C.Server.Host}).Info("Starting server")
	if err = r.Run(fmt.Sprintf("%s:%d", conf.C.Server.Host, conf.C.Server.Port)); err != nil {
		return err
	}
	return nil
}
