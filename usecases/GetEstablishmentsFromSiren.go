package usecases

import (
	"errors"

	"github.com/jclebreton/opensirene/domain"
)

type GetEstablishmentsFromSirenRequest struct {
	Siren  string
	Limit  string
	Offset string
}

// GetEstablishmentsFromSiren returns the requested establishments
func (i *Interactor) GetEstablishmentsFromSiren(r GetEstablishmentsFromSirenRequest) (*[]domain.Establishment, error) {
	return r.findEstablishmentsFromSiren(i)
}

func (r *GetEstablishmentsFromSirenRequest) findEstablishmentsFromSiren(i *Interactor) (*[]domain.Establishment, error) {
	ee, err := i.EnterprisesRW.FindEstablishmentsFromSiren(r.Siren, r.Limit, r.Offset)
	if err != nil {
		return nil, err
	}

	if ee == nil {
		return nil, errors.New("nothing found")
	}

	return ee, nil
}
