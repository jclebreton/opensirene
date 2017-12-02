package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jclebreton/opensirene/usecases"
	"github.com/sirupsen/logrus"
)

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
