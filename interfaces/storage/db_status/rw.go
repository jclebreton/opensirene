package db_status

import (
	"github.com/jackc/pgx"
	"github.com/jclebreton/opensirene/domain"
)

type RW struct {
	PgxClient *pgx.ConnPool
}

func (rw *RW) FindDatabaseStatus() ([]domain.UpdateFileStatus, error) {
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

// GetSuccessList lists the successful updates in the database
// and returns the slice of file names which were successfully imported.
// Returns an empty slice otherwise.
func (rw *RW) GetSuccessList() ([]string, error) {
	rows, err := rw.PgxClient.Query("SELECT filename FROM history WHERE is_success=true")
	if err != nil {
		return nil, err
	}

	var result []string
	for rows.Next() {
		var filename string
		if err := rows.Scan(&filename); err != nil {
			return nil, err
		}
		result = append(result, filename)
	}

	return result, nil
}
