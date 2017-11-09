package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jclebreton/opensirene/interfaces/json"
)

func (h HttpGateway) GetHistories(c *gin.Context) {
	hh, err := h.i.FindHistories()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, json.GetHistoriesRespFormatFormat(hh))
	return
}
