package db_status

import "github.com/jackc/pgx"

type RW struct {
	PgxClient *pgx.ConnPool
}

// GetSuccessList the slice of file names which were successfully imported.
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
