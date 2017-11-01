package views

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jclebreton/opensirene/api/models"
	"github.com/jclebreton/opensirene/database"
)

// GetHistory is in charge of querying the database to get database history
func GetHistory(c *gin.Context) {
	var h models.Histories

	if database.DB.Find(&h).RecordNotFound() {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, h)
}
