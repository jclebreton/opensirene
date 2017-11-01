package models

import "time"

// History is a struct mapping history sql table
type History struct {
	ID        int32     `gorm:"primary_key,column:id"`
	Datetime  time.Time `gorm:"column:datetime"`
	Action    string    `gorm:"column:action"`
	IsSuccess bool      `gorm:"column:is_success"`
	Filename  string    `gorm:"column:filename"`
	Msg       string    `gorm:"column:msg"`
}

// Histories is a slice of History
type Histories []History

func (Histories) TableName() string {
	return "history"
}
