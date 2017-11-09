package usecases

import "github.com/jclebreton/opensirene/domain"

type JsonW interface {
	FormatGetDBStatusResp(hh []domain.UpdateFileStatus) interface{}
	FormatGetEnterpriseFromSiretResp(e domain.Establishment) interface{}
	FormatGetEstablishmentsFromSirenResp(es []domain.Establishment) interface{}
}

type DBStatusRW interface {
	FindDatabaseStatus() ([]domain.UpdateFileStatus, error)
	GetSuccessList() ([]string, error)
}

type EnterprisesRW interface {
	FindEnterpriseBySiret(siret string) (*domain.Establishment, error)
	FindEstablishmentsFromSiren(siren, limit, offset string) ([]domain.Establishment, error)
}
