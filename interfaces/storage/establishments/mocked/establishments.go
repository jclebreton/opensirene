package mocked

import "github.com/jclebreton/opensirene/domain"

type RW struct {
	FindEnterpriseBySiretRet
	FindEstablishmentsFromSirenRet
}

func (rw RW) FindEnterpriseBySiret(siret string) (*domain.Establishment, error) {
	return rw.FindEnterpriseBySiretRet.DBStatus, rw.FindEnterpriseBySiretRet.Err
}

type FindEnterpriseBySiretRet struct {
	DBStatus *domain.Establishment
	Err      error
}

func (rw RW) FindEstablishmentsFromSiren(siren, limit, offset string) (*[]domain.Establishment, error) {
	return rw.FindEstablishmentsFromSirenRet.DBStatus, rw.FindEstablishmentsFromSirenRet.Err
}

type FindEstablishmentsFromSirenRet struct {
	DBStatus *[]domain.Establishment
	Err      error
}
