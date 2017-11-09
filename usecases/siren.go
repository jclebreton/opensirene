package usecases

import "github.com/jclebreton/opensirene/domain"

func (i Interactor) GetHistories() ([]domain.History, error) {
	hh, err := i.HistoryRW.FindHistories()
	if err != nil {
		return nil, err
	}
	return hh, nil
}
