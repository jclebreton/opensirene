package logic

import (
	"github.com/jclebreton/opensirene/database"
	"github.com/pkg/errors"
)

type track struct {
	Database database.PgxClient
}

type Tracker interface {
	Save(action, msg, filename string, isSuccess bool) error
}

func NewTracker(db database.PgxClient) *track {
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
