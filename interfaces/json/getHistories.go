package json

import (
	"time"

	"github.com/jclebreton/opensirene/domain"
)

//
// THE STANDARD VERSION
//
type formatGetHistoriesRespStd struct{}

type historyJSON struct {
	ID        int32     `json:"id"`
	Datetime  time.Time `json:"datetime"`
	Filename  string    `json:"filename"`
	IsSuccess bool      `json:"is_success"`
	Err       string    `json:"err"`
}

func (jw formatGetHistoriesRespStd) FormatGetHistoriesResp(hh []domain.History) interface{} {
	hhJSON := []historyJSON{}
	for _, h := range hh {
		hhJSON = append(hhJSON, historyJSON{
			ID:        h.ID,
			Datetime:  h.Datetime,
			Filename:  h.Filename,
			IsSuccess: h.IsSuccess,
			Err:       h.Err,
		})
	}
	return hhJSON
}

//
// THE SIMPLE DATES VERSION
//
type formatGetHistoriesRespSimpleDates struct{}

type historySimpleDatesJSON struct {
	ID        int32  `json:"id"`
	Datetime  string `json:"datetime"`
	Filename  string `json:"filename"`
	IsSuccess bool   `json:"is_success"`
	Err       string `json:"err"`
}

func (jw formatGetHistoriesRespSimpleDates) FormatGetHistoriesResp(hh []domain.History) interface{} {
	hhJSON := []historySimpleDatesJSON{}
	for _, h := range hh {
		hhJSON = append(hhJSON, historySimpleDatesJSON{
			ID:        h.ID,
			Datetime:  h.Datetime.Format("2006-01-02:15h04"),
			Filename:  h.Filename,
			IsSuccess: h.IsSuccess,
			Err:       h.Err,
		})
	}
	return hhJSON
}
