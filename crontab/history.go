package crontab

import (
	"github.com/jackc/pgx"
	"github.com/pkg/errors"
)

// LogImport save to db the various milestones
func LogImport(db *pgx.ConnPool, action, msg, filename string, isSuccess bool) error {
	var err error

	_, err = db.Exec(
		"INSERT INTO history (action, is_success, filename, msg) VALUES ($1, $2, $3, $4)",
		action,
		isSuccess,
		filename,
		msg,
	)

	if err != nil {
		return errors.Wrap(err, "couldn't log sql transaction")
	}

	return nil
}
