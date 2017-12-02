package usecases

import (
	"errors"

	"github.com/jclebreton/opensirene/domain"
)

type GetEnterpriseFromSiretRequest struct {
	Siret string
}

// GetEnterpriseFromSiret returns the requested enterprise
func (i *Interactor) GetEnterpriseFromSiret(r GetEnterpriseFromSiretRequest) (*domain.Establishment, error) {
	return r.findEnterpriseFromSiret(i)
}

func (r *GetEnterpriseFromSiretRequest) findEnterpriseFromSiret(i *Interactor) (*domain.Establishment, error) {
	e, err := i.EnterprisesRW.FindEnterpriseBySiret(r.Siret)
	if err != nil {
		return nil, err
	}

	if e == nil {
		return nil, errors.New("nothing found")
	}

	return e, nil
}
