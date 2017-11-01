package views

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/jclebreton/opensirene/api/models"
	"github.com/jclebreton/opensirene/database"
)

// GetSiren is in charge of querying the database to get the specific records
// for a single Siren given in the query
func GetSiren(c *gin.Context) {
	var err error
	var es models.Enterprises
	limit := -1
	offset := -1
	siren := c.Param("id")
	lim := c.DefaultQuery("limit", "")
	off := c.DefaultQuery("offset", "")

	if lim != "" {
		if limit, err = strconv.Atoi(lim); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "'limit' query parameter isn't an integer"})
			return
		}
	}
	if off != "" {
		if offset, err = strconv.Atoi(off); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "'offset' query parameter isn't an integer"})
			return
		}
	}

	res := database.DB.Limit(limit).Offset(offset).Order("siret ASC").Find(&es, models.Enterprise{Siren: siren})
	if res.RecordNotFound() || len(es) == 0 {
		c.Status(http.StatusNotFound)
		return
	}
	if res.Error != nil {
		c.Status(http.StatusInternalServerError)
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
