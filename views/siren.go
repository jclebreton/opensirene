package views

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jclebreton/opensirene/database"
	"github.com/jclebreton/opensirene/models"
)

// GetSiren is in charge of querying the database to get the specific records
// for a single Siren given in the query
func GetSiren(c *gin.Context) {
	var es models.Enterprises

	siren := c.Param("id")

	if database.DB.Find(&es, models.Enterprise{Siren: siren}).RecordNotFound() || len(es) == 0 {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, es)
}

// GetSiret is in charge of querying the database to get the specific enterprise
// record for a single Siret given in the query
func GetSiret(c *gin.Context) {
	var e models.Enterprise

	siret := c.Param("id")

	if database.DB.Find(&e, models.Enterprise{Siret: siret}).RecordNotFound() {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, e)
}
