package monitoring

import "github.com/jclebreton/opensirene/domain"

type MonitoringRW interface {
	GetHealth() *domain.Health
	GetPing() domain.Pong
	GetDBStatus() ([]domain.UpdateFileStatus, error)
}
