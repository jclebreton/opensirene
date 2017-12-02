package json

import (
	"time"

	"github.com/jclebreton/opensirene/domain"
)

type formatGetDBStatusResp struct{}

type UpdateFileStatusJSON struct {
	ID        int32     `json:"id"`
	Datetime  time.Time `json:"datetime"`
	Filename  string    `json:"filename"`
	IsSuccess bool      `json:"is_success"`
	Err       string    `json:"err"`
}

func (jw *formatGetDBStatusResp) FormatGetDBStatusResp(hh []domain.UpdateFileStatus) interface{} {
	var hhJSON []UpdateFileStatusJSON
	for _, h := range hh {
		hhJSON = append(hhJSON, UpdateFileStatusJSON{
			ID:        h.ID,
			Datetime:  h.Datetime,
			Filename:  h.Filename,
			IsSuccess: h.IsSuccess,
			Err:       h.Err,
		})
	}
	return hhJSON
}
