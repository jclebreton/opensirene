package usecases

import "github.com/jclebreton/opensirene/domain"

type HistoryRW interface {
	FindHistories() ([]domain.History, error)
}

type JsonW interface {
	FormatGetHistoriesResp(hh []domain.History) interface{}
}
