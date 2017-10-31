package models

import (
	"time"

	"github.com/jclebreton/opensirene/database"
)

// History is a struct mapping history sql table
type History struct {
	ID        int32     `gorm:"primary_key,column:id"`
	Datetime  time.Time `gorm:"column:datetime"`
	Action    string    `gorm:"column:action"`
	isSuccess bool      `gorm:"column:is_success"`
	Filename  string    `gorm:"column:filename"`
	Msg       string    `gorm:"column:msg"`
}

func GetSuccessfulUpdateList() []string {
	var sh []History
	if database.DB.Find(&sh, History{isSuccess: true}).RecordNotFound() {
		return []string{}
	}

	var r []string
	for _, h := range sh {
		r = append(r, h.Filename)
	}
	return r
}
