package crontab

import (
	"errors"

	"github.com/jclebreton/opensirene/database"
)

type mutex struct {
	Database database.PgxClient
}

type Mutexer interface {
	Lock() error
	Unlock() error
}

func NewMutex(db database.PgxClient) *mutex {
	return &mutex{Database: db}
}

// TryLock set a mutex for database write
func (m *mutex) Lock() error {
	var result bool
	err := m.Database.Conn.QueryRow("SELECT pg_try_advisory_lock(42)").Scan(&result)
	if err != nil {
		return err
	}
	if !result {
		return errors.New("database mutex is already taken")
	}
	return nil
}

// Unlock remove the current database mutex
func (m *mutex) Unlock() error {
	var result bool
	err := m.Database.Conn.QueryRow("SELECT pg_advisory_unlock(42)").Scan(&result)
	if err != nil {
		return err
	}
	if !result {
		return errors.New("unable to release database mutex")
	}
	return nil
}
