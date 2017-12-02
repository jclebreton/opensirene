package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jclebreton/opensirene/usecases"
	"github.com/sirupsen/logrus"
)

func (h HttpGateway) Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
	return
}

func (h HttpGateway) Health(c *gin.Context) {
	health, err := h.i.MonitoringR.GetHealth()
	if err != nil {
		logrus.WithError(err).Error("error")
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, h.i.JsonW.FormatHealthResp(health))
	return
}

func (h HttpGateway) GetDBStatus(c *gin.Context) {
	hh, err := h.i.GetDBStatus(usecases.GetDBStatusRequest{})
	if err != nil {
		logrus.WithError(err).Error("error")
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, h.i.JsonW.FormatGetDBStatusResp(hh))
	return
}
