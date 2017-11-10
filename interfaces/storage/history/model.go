package history

import (
	"time"

	"github.com/jclebreton/opensirene/domain"
)

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
func (History) TableName() string {
	return "history"
}

func (h History) ToUC() *domain.History {
	return &domain.History{
		ID:        h.ID,
		Datetime:  h.Datetime,
		Filename:  h.Filename,
		IsSuccess: h.IsSuccess,
		Err:       h.Err,
	}
}
