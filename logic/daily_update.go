package logic

import (
	"github.com/jclebreton/opensirene/api/models"
	"github.com/jclebreton/opensirene/database"
	"github.com/jclebreton/opensirene/opendata/gouv_sirene"
	"github.com/sirupsen/logrus"
)

func GetSuccessfulUpdateList() []string {
	var sh []models.History
	if database.DB.Find(&sh, models.History{IsSuccess: true}).RecordNotFound() {
		return []string{}
	}

	var r []string
	for _, h := range sh {
		r = append(r, h.Filename)
	}
	return r
}

// Daily is the cron task that runs every few hours to get and apply the latest
// updates
func Daily() {
	var err error
	var sfs gouv_sirene.RemoteFiles

	if sfs, err = gouv_sirene.GrabLatestFull(); err != nil {
		logrus.WithError(err).Error("Could not download latest")
		return
	}

	sfs = sfs.Diff(GetSuccessfulUpdateList())

	if err = Import(sfs); err != nil {
		logrus.WithError(err).Error("Could not download latest")
		return
	}
}
