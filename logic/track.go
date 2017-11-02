package logic

import (
	"github.com/pkg/errors"

	"github.com/jclebreton/opensirene/database"
)

type track struct {
	Database database.PgxClient
}

// Tracker is an interface used to track success and errors
type Tracker interface {
	Save(action, msg, filename string, isSuccess bool) error
}

// newTracker returns a new track
func newTracker(db database.PgxClient) *track {
	return &track{Database: db}
}

// Save logs to database the milestones
func (t *track) Save(action, msg, filename string, isSuccess bool) error {
	var err error

	_, err = t.Database.Conn.Exec(
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

// Truncate will erase all logs
func (t *track) Truncate() error {
	var err error

	_, err = t.Database.Conn.Exec("TRUNCATE TABLE history")

	if err != nil {
		return errors.Wrap(err, "couldn't truncate history table")
	}

	return nil
}
