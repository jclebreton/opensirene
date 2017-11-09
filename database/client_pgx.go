package database

import (
	"github.com/jackc/pgx"

	"github.com/jclebreton/opensirene/conf"
)

// PgxClient represents a pgx client
type PgxClient struct {
	Conn *pgx.ConnPool
}

// NewImportClient creates a new PgxClient from the configuration
func NewPgxClientClient() (*pgx.ConnPool, error) {
	connPoolConfig := pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Database: conf.C.Database.Name,
			Host:     conf.C.Database.Host,
			Port:     uint16(conf.C.Database.Port),
			User:     conf.C.Database.User,
			Password: conf.C.Database.Password,
		},
		MaxConnections: 5,
		AfterConnect: func(conn *pgx.Conn) error {
			_, err := conn.Exec("SET CLIENT_ENCODING TO 'UTF-8'")
			return err
		},
	}

	return pgx.NewConnPool(connPoolConfig)
}
