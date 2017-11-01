package crontab

import (
	"github.com/jclebreton/opensirene/database"
	"github.com/pkg/errors"
)

type Track struct {
	Database database.PgxClient
}

type Tracker interface {
	Save(action, msg, filename string, isSuccess bool) error
}

func NewTracker(db database.PgxClient) *mutex {
	return &mutex{Database: db}
}

// Save logs to database the milestones
func (t *Track) Save(action, msg, filename string, isSuccess bool) error {
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
