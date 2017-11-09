package json

import (
	"time"

	"github.com/jclebreton/opensirene/domain"
)

type HistoryJSON struct {
	ID        int32     `json:"id"`
	Datetime  time.Time `json:"datetime"`
	Filename  string    `json:"filename"`
	IsSuccess bool      `json:"is_success"`
	Err       string    `json:"err"`
}

func newFromHistory(h domain.History) HistoryJSON {
	return HistoryJSON{
		ID:        h.ID,
		Datetime:  h.Datetime,
		Filename:  h.Filename,
		IsSuccess: h.IsSuccess,
		Err:       h.Err,
	}
}

func GetHistoriesRespFormatFormat(hh []domain.History) []HistoryJSON {
	hhJSON := []HistoryJSON{}
	for _, h := range hh {
		hhJSON = append(hhJSON, newFromHistory(h))
	}
	return hhJSON
}
