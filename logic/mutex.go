package logic

import (
	"errors"

	"github.com/jclebreton/opensirene/database"
)

// mutex is an implementation of the Mutexer interface
type mutex struct {
	Database database.PgxClient
}

// Mutexer is an interface to use the mutex logic in postgreSQL
type Mutexer interface {
	Lock() error
	Unlock() error
}

// newMutex returns a new Mutex
func newMutex(db database.PgxClient) *mutex {
	return &mutex{Database: db}
}

// Lock set a mutex for database write
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
