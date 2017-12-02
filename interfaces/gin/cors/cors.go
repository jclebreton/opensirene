package cors

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type StandardCors struct {
	permissive     bool
	allowedOrigins []string
}

func NewStandardCors(permissive bool, allowedOrigins []string) StandardCors {
	return StandardCors{
		permissive, allowedOrigins,
	}
}

func (sCors StandardCors) SetCors() gin.HandlerFunc {
	cc := cors.Config{
		AllowMethods:  []string{"POST", "GET", "OPTIONS", "PUT", "DELETE", "PATCH"},
		AllowHeaders:  []string{"DNT", "X-CustomHeader", "Keep-Alive", "User-Agent", "X-Requested-With", "If-Modified-Since", "Cache-Control", "Content-Type", "Content-Range", "Range", "Authorization"},
		ExposeHeaders: []string{"DNT", "X-CustomHeader", "Keep-Alive", "User-Agent", "X-Requested-With", "If-Modified-Since", "Cache-Control", "Content-Type", "Content-Range", "Range", "Authorization"},
	}
	if sCors.permissive {
		cc.AllowAllOrigins = true
	} else {
		cc.AllowOrigins = sCors.allowedOrigins
	}
	return cors.New(cc)
}
