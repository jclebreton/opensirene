package models

import "time"

// History is a struct mapping history sql table
type History struct {
	ID        int32     `gorm:"primary_key,column:id"`
	Datetime  time.Time `gorm:"column:datetime"`
	Filename  string    `gorm:"column:filename"`
	IsSuccess bool      `gorm:"column:is_success"`
	Err       string    `gorm:"column:err"`
}

// Histories is a slice of History
type Histories []History

// TableName overrides the table name calculated by gorm
func (Histories) TableName() string {
	return "history"
}
