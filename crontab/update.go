package crontab

import (
	"github.com/jackc/pgx"
	"github.com/jclebreton/opensirene/api/models"
	"github.com/jclebreton/opensirene/database"
	"github.com/pkg/errors"
)

func GetSuccessfulUpdateList() []string {
	var sh []models.History
	if database.DB.Find(&sh, models.History{IsSuccess: true}).RecordNotFound() {
		return []string{}
	}

	var r []string
	for _, h := range sh {
		r = append(r, h.Filename)
	}
	return r
}

// Update update stock table from daily update file
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

// TryLock set a mutex for database write
func TryLock() error {
	var result bool
	err := database.ImportClient.Conn.QueryRow("SELECT pg_try_advisory_lock(42)").Scan(&result)
	if err != nil {
		return err
	}
	if !result {
		return errors.New("mutex is already taken")
	}
	return nil
}

// Unlock remove the current database mutex
func Unlock() error {
	var result bool
	err := database.ImportClient.Conn.QueryRow("SELECT pg_advisory_unlock(42)").Scan(&result)
	if err != nil {
		return err
	}
	if !result {
		return errors.New("unable to release the mutex")
	}
	return nil
}
