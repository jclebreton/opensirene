package usecases

import (
	"github.com/jclebreton/opensirene/domain"
	"github.com/pkg/errors"
)

type GetDBStatusRequest struct{}

// GetDBStatus returns the list of updates applied in the database
func (i *Interactor) GetDBStatus(r GetDBStatusRequest) ([]domain.UpdateFileStatus, error) {
	return r.findDBSatus(i)
}

func (r *GetDBStatusRequest) findDBSatus(i *Interactor) ([]domain.UpdateFileStatus, error) {
	hh, err := i.DBStatusRW.FindDatabaseStatus()
	if err != nil {
		return nil, err
	}

	if hh == nil || len(hh) == 0 {
		return nil, errors.New("nothing found")
	}

	return hh, nil
}
