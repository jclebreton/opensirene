package db_status

import (
	"time"

	"github.com/jclebreton/opensirene/domain"
)

// DBStatus is a struct mapping history sql table
type UpdateStatus struct {
	ID        int32     `gorm:"primary_key,column:id"`
	Datetime  time.Time `gorm:"column:datetime"`
	Filename  string    `gorm:"column:filename"`
	IsSuccess bool      `gorm:"column:is_success"`
	Err       string    `gorm:"column:err"`
}

// TableName overrides the table name calculated by gorm
func (UpdateStatus) TableName() string {
	return "history"
}

func (h *UpdateStatus) ToUC() *domain.UpdateFileStatus {
	return &domain.UpdateFileStatus{
		ID:        h.ID,
		Datetime:  h.Datetime,
		Filename:  h.Filename,
		IsSuccess: h.IsSuccess,
		Err:       h.Err,
	}
}
