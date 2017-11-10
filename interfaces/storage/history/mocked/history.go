package mocked

import "github.com/jclebreton/opensirene/domain"

type RW struct {
	FindHistoriesRet
}

func (rw RW) FindHistories() ([]domain.History, error) {
	return rw.FindHistoriesRet.Histories, rw.FindHistoriesRet.Err
}

type FindHistoriesRet struct {
	Histories []domain.History
	Err       error
}
