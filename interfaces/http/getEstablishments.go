package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jclebreton/opensirene/usecases"
	"github.com/sirupsen/logrus"
)

func (h HttpGateway) GetEstablishmentsFromSiren(c *gin.Context) {
	r := usecases.GetEstablishmentsFromSirenRequest{
		Siren:  c.Param("id"),
		Offset: c.Param("offset"),
		Limit:  c.Param("limit"),
	}
	es, err := h.i.GetEstablishmentsFromSiren(r)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, h.i.JsonW.FormatGetEstablishmentsFromSirenResp(es))
	return
}

func (h HttpGateway) GetEnterpriseFromSiret(c *gin.Context) {
	r := usecases.GetEnterpriseFromSiretRequest{
		Siret: c.Param("id"),
	}
	e, err := h.i.GetEnterpriseFromSiret(r)
	if err != nil {
		logrus.WithError(err).Info("SIRET")
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, h.i.JsonW.FormatGetEnterpriseFromSiretResp(*e))
	return
}
