package logic

import (
	"github.com/jclebreton/opensirene/api/models"
	"github.com/jclebreton/opensirene/database"
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
