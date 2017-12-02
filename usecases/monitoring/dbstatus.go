package monitoring

import (
	"github.com/jclebreton/opensirene/domain"
	"github.com/pkg/errors"
)

type GetDBStatusRequest struct{}

// GetDBStatus returns the list of updates applied in the database
func (i *Interactor) GetDBStatus(r GetDBStatusRequest) ([]domain.UpdateFileStatus, error) {
	return r.getDBStatus(i)
}

func (r *GetDBStatusRequest) getDBStatus(i *Interactor) ([]domain.UpdateFileStatus, error) {
	hh, err := i.MonitoringRW.GetDBStatus()
	if err != nil {
		return nil, err
	}

	if hh == nil || len(hh) == 0 {
		return nil, errors.New("nothing found")
	}

	return hh, nil
}
