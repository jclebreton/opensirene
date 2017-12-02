package usecases

import "github.com/jclebreton/opensirene/domain"

type JsonW interface {
	FormatHealthResp(hh *domain.Health) interface{}
	FormatGetDBStatusResp(hh []domain.UpdateFileStatus) interface{}
	FormatGetEnterpriseFromSiretResp(e *domain.Establishment) interface{}
	FormatGetEstablishmentsFromSirenResp(es *[]domain.Establishment) interface{}
}

type MonitoringR interface {
	GetHealth() (*domain.Health, error)
	GetPing() (domain.Pong, error)
}

type DBStatusRW interface {
	GetSuccessList() ([]string, error)
}

type EnterprisesRW interface {
	FindEnterpriseBySiret(siret string) (*domain.Establishment, error)
	FindEstablishmentsFromSiren(siren, limit, offset string) (*[]domain.Establishment, error)
}

type SireneR interface {
	GetRemoteFiles(downloadPath string) ([]domain.RemoteFile, error)
}
