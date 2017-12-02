package monitoring

import (
	"github.com/Depado/ginprom"
	"github.com/gin-gonic/gin"
)

type PrometheusStd struct {
	prefix string
	engine *gin.Engine
}

func NewPrometheus(prefix string, engine *gin.Engine) PrometheusStd {
	return PrometheusStd{prefix, engine}
}

func (p PrometheusStd) Monitor() gin.HandlerFunc {
	return ginprom.New(ginprom.Subsystem(p.prefix), ginprom.Engine(p.engine)).Instrument()
}
