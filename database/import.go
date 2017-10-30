package database

import (
	"github.com/jackc/pgx"

	"github.com/jclebreton/opensirene/conf"
	"github.com/pkg/errors"
)

type PgxClient struct {
	Conn *pgx.ConnPool
}

var ImportClient PgxClient

func NewImportClient() (*PgxClient, error) {
	connectConfig := &pgx.ConnConfig{
		Database: conf.C.Database.Name,
		Host:     conf.C.Database.Host,
		Port:     uint16(conf.C.Database.Port),
		User:     conf.C.Database.User,
		Password: conf.C.Database.Password,
	}

	connPoolConfig := pgx.ConnPoolConfig{
		ConnConfig:     *connectConfig,
		MaxConnections: 5,
		AfterConnect: func(conn *pgx.Conn) error {
			_, err := conn.Exec("SET CLIENT_ENCODING TO 'UTF-8'")
			return err
		},
	}

	pool, err := pgx.NewConnPool(connPoolConfig)
	if err != nil {
		return nil, err
	}

	return &PgxClient{Conn: pool}, nil
}

func InitImportClient() error {
	connectConfig := &pgx.ConnConfig{
		Database: conf.C.Database.Name,
		Host:     conf.C.Database.Host,
		Port:     uint16(conf.C.Database.Port),
		User:     conf.C.Database.User,
		Password: conf.C.Database.Password,
	}

	connPoolConfig := pgx.ConnPoolConfig{
		ConnConfig:     *connectConfig,
		MaxConnections: 5,
		AfterConnect: func(conn *pgx.Conn) error {
			_, err := conn.Exec("SET CLIENT_ENCODING TO 'UTF-8'")
			return err
		},
	}

	pool, err := pgx.NewConnPool(connPoolConfig)
	if err != nil {
		return err
	}

	ImportClient = PgxClient{
		Conn: pool,
	}

	return nil
}

// TryLock set a mutex for database write
func (c *PgxClient) TryLock() error {
	var result bool
	err := c.Conn.QueryRow("SELECT pg_try_advisory_lock(42)").Scan(&result)
	if err != nil {
		return err
	}
	if !result {
		return errors.New("mutex is already taken")
	}
	return nil
}

// Unlock remove the current mutex
func (c *PgxClient) Unlock() error {
	var result bool
	err := c.Conn.QueryRow("SELECT pg_advisory_unlock(42)").Scan(&result)
	if err != nil {
		return err
	}
	if !result {
		return errors.New("unable to release the mutex")
	}
	return nil
}
