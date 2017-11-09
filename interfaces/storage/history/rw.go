package history

import (
	"github.com/jclebreton/opensirene/domain"
	"github.com/jinzhu/gorm"
)

type RW struct {
	GormClient *gorm.DB
}

func (rw RW) FindHistories() ([]domain.History, error) {
	hh := Histories{}
	if rw.GormClient.Find(&hh).RecordNotFound() {
		return nil, rw.GormClient.Error
	}

	dHistory := []domain.History{}
	for _, h := range hh {
		dHistory = append(dHistory, *h.ToUC())
	}

	return dHistory, nil
}
