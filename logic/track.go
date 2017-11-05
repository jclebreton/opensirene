package logic

import (
	"github.com/pkg/errors"

	"github.com/jclebreton/opensirene/database"
)

type track struct {
	pgxClient *database.PgxClient
}

// Tracker is an interface used to track success and errors
type Tracker interface {
	Save(filename string, isSuccess bool, err string) error
	Truncate() error
}

// newTracker returns a new track
func NewTracker(pgxClient *database.PgxClient) *track {
	return &track{pgxClient: pgxClient}
}

// Save logs to database the milestones
func (t *track) Save(filename string, isSuccess bool, err string) error {
	if _, err := t.pgxClient.Conn.Exec(
		"INSERT INTO history (filename, is_success, err) VALUES ($1, $2, $3)",
		filename,
		isSuccess,
		err,
	); err != nil {
		return errors.Wrap(err, "couldn't log sql transaction")
	}

	return nil
}

// Truncate will erase all logs
func (t *track) Truncate() error {
	if _, err := t.pgxClient.Conn.Exec("TRUNCATE TABLE history"); err != nil {
		return errors.Wrap(err, "couldn't truncate history table")
	}
	return nil
}
