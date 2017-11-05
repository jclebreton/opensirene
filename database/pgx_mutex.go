package database

import "errors"

// mutex is an implementation of the Mutexer interface
type pgxMutex struct {
	pgxClient *PgxClient
}

// newMutex returns a new Mutex
func NewPgxMutex(db *PgxClient) *pgxMutex {
	return &pgxMutex{pgxClient: db}
}

// Lock set a mutex for database write
func (m *pgxMutex) Lock() error {
	var result bool
	err := m.pgxClient.Conn.QueryRow("SELECT pg_try_advisory_lock(42)").Scan(&result)
	if err != nil {
		return err
	}
	if !result {
		return errors.New("database mutex is already taken")
	}
	return nil
}

// Unlock remove the current database mutex
func (m *pgxMutex) Unlock() error {
	var result bool
	err := m.pgxClient.Conn.QueryRow("SELECT pg_advisory_unlock(42)").Scan(&result)
	if err != nil {
		return err
	}
	if !result {
		return errors.New("unable to release database mutex")
	}
	return nil
}
