package monitoring

import (
	"github.com/jackc/pgx"
	"github.com/jclebreton/opensirene/domain"
)

type MonitoringRW struct {
	PgxClient *pgx.ConnPool
	Version   string
}

func (rw *MonitoringRW) GetHealth() *domain.Health {
	return &domain.Health{
		Name:    "opensirene",
		Version: rw.Version,
	}
}

func (rw *MonitoringRW) GetPing() domain.Pong {
	return "pong"
}

func (rw *MonitoringRW) GetDBStatus() ([]domain.UpdateFileStatus, error) {
	rows, err := rw.PgxClient.Query("SELECT id, datetime, filename, is_success, err FROM history")
	if err != nil {
		return nil, err
	}

	var result []domain.UpdateFileStatus
	for rows.Next() {
		r := domain.UpdateFileStatus{}
		if err := rows.Scan(&r.ID, &r.Datetime, &r.Filename, &r.IsSuccess, &r.Err); err != nil {
			return nil, err
		}
		result = append(result, r)
	}

	return result, nil
}
