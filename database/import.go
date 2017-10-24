package database

import (
	"github.com/jackc/pgx"

	"github.com/Depado/lightsiren/conf"
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
