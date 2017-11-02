package database

import (
	"github.com/jackc/pgx"

	"github.com/jclebreton/opensirene/conf"
)

// PgxClient represents a pgx client
type PgxClient struct {
	Conn *pgx.ConnPool
}

// ImportClient is the main exported client
var ImportClient PgxClient

// NewImportClient creates a new PgxClient from the configuration
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

// InitImportClient intializes the main client
func InitImportClient() error {
	client, err := NewImportClient()
	if err != nil {
		return err
	}
	ImportClient = *client
	return nil
}
