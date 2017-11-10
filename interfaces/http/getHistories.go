package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jclebreton/opensirene/usecases"
)

func (h HttpGateway) GetHistories(c *gin.Context) {
	// here you'd setup everything so the interactor method just has to deal with abstract (domain/usecases) objects

	hh, err := h.i.GetHistories(usecases.GetHistoriesRequest{})
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, h.i.JsonW.FormatGetHistoriesResp(hh))
	return
}
