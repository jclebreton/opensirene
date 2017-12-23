package views

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jclebreton/opensirene/api/models"
	"github.com/jclebreton/opensirene/conf"
	"github.com/jclebreton/opensirene/database"
	"github.com/jclebreton/opensirene/opendata/gouvfr/sirene"
	"github.com/sirupsen/logrus"
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

	// check database connexion
	dbErr := ""
	dbStatus := "OK"
	if _, err := database.NewImportClient(); err != nil {
		logrus.WithError(err).Error("Couldn't check database connexion")
		dbStatus = "KO"
		dbErr = err.Error()
	}

	//check gouv API
	gouvErr := ""
	gouvStatus := "OK"
	if _, err := http.Head(sirene.DatasetEndpoint + sirene.SirenID); err != nil {
		logrus.WithError(err).Error("Couldn't check gouv API connexion")
		gouvStatus = "KO"
		gouvErr = err.Error()
	}

	//check sirene files
	filesURL := "http://files.data.gouv.fr/sirene/"
	filesErr := ""
	filesStatus := "OK"
	if _, err := http.Head(filesURL); err != nil {
		logrus.WithError(err).Error("Couldn't check sirene connexion")
		filesStatus = "KO"
		filesErr = err.Error()
	}

	h := models.Health{
		Name:      "opensirene",
		Version:   v.Version,
		BuildDate: v.BuildDate,
		Dependencies: map[string]models.Dependency{
			conf.C.Database.Name: {
				Name:   conf.C.Database.Host,
				Status: dbStatus,
				Error:  dbErr,
			},
			"www.data.gouv.fr": {
				Name:   sirene.DatasetEndpoint + sirene.SirenID,
				Status: gouvStatus,
				Error:  gouvErr,
			},
			"files.data.gouv.fr": {
				Name:   filesURL,
				Status: filesStatus,
				Error:  filesErr,
			},
		},
	}

	logrus.WithField("health", h).Info()

	c.JSON(http.StatusOK, h)
}
