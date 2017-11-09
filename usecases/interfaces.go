package usecases

import "github.com/jclebreton/opensirene/domain"

type HistoryRW interface {
	FindHistories() ([]domain.History, error)
}
