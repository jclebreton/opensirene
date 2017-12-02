package mocked

import "github.com/jclebreton/opensirene/domain"

type MonitoringRW struct {
	GetDatabaseStatusRet
	GetPingRet
	GetHealthRet
}

type GetDatabaseStatusRet struct {
	DBStatus []domain.UpdateFileStatus
	Err      error
}

type GetPingRet struct {
	Value domain.Pong
}

type GetHealthRet struct {
	Health *domain.Health
}

func (rw MonitoringRW) GetDatabaseStatus() ([]domain.UpdateFileStatus, error) {
	return rw.GetDatabaseStatusRet.DBStatus, rw.GetDatabaseStatusRet.Err
}

func (rw MonitoringRW) GetPing() domain.Pong {
	return rw.GetPingRet.Value
}

func (rw *MonitoringRW) GetHealth(version string) *domain.Health {
	return rw.GetHealthRet.Health
}
