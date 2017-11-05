package views

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/jclebreton/opensirene/api/models"
)

// GetSiren is in charge of querying the database to get the specific records
// for a single Siren given in the query
func (v *ViewsContext) GetSiren(c *gin.Context) {
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

	res := v.GormClient.Limit(limit).Offset(offset).Order("nic ASC").Find(&es, models.Enterprise{Siren: siren})
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
func (v *ViewsContext) GetSiret(c *gin.Context) {
	var e models.Enterprise

	siret := c.Param("id")
	if len(siret) != 14 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "not a valid siret: must be 14-digit number"})
		return
	}

	if v.GormClient.Find(&e, models.Enterprise{Siren: siret[0:9], Nic: siret[9:14]}).RecordNotFound() {
		c.Status(http.StatusNotFound)
		return
	}

	e.Siret = e.Siren + e.Nic

	c.JSON(http.StatusOK, e)
}
