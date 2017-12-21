package views

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jclebreton/opensirene/api/models"
	"github.schibsted.io/opensirene/conf"
)

// GetHistory is in charge of querying the database to get database history
func (v *ViewsContext) GetHistory(c *gin.Context) {
	var h models.Histories

	if v.GormClient.Find(&h).RecordNotFound() {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, h)
}

// GetPing is a monitoring endpoint
func (v *ViewsContext) GetPing(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

// GetHealth is a monitoring endpoint
func (v *ViewsContext) GetHealth(c *gin.Context) {
	h := models.Health{
		Name:    "opensirene",
		Version: v.Version,
		Dependencies: map[string]models.Dependency{
			conf.C.Database.Host: models.Dependency{
				Name:   conf.C.Database.Name,
				Status: "Unknown",
				Error:  "",
			},
			"www.data.gouv.fr": models.Dependency{
				Name:   "www.data.gouv.fr",
				Status: "Unknown",
				Error:  "",
			},
		},
	}
	c.JSON(http.StatusOK, h)
}
