package models

import "time"

// History is a struct mapping history sql table
type History struct {
	ID        int32     `gorm:"primary_key,column:id"`
	IsSuccess bool      `gorm:"column:is_success"`
	Action    string    `gorm:"column:action"`
	Filename  string    `gorm:"column:filename"`
	Msg       string    `gorm:"column:msg"`
	Datetime  time.Time `gorm:"column:datetime"`
}

// Histories is a slice of History
type Histories []History

// TableName overrides the table name calculated by gorm
func (Histories) TableName() string {
	return "history"
}
