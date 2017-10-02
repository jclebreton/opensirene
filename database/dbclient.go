package database

import "github.com/jackc/pgx"

type DBClient struct {
	conn *pgx.ConnPool
}

func InitDBClient() (*DBClient, error) {
	connectConfig := &pgx.ConnConfig{
		Database: "opensirenedb",
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "",
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

	client := &DBClient{}
	client.conn = pool

	return client, nil
}

func (client *DBClient) ImportCSVFile(table string, columns []string, source pgx.CopyFromSource) (int, error) {
	copyCount, err := client.conn.CopyFrom(pgx.Identifier{table}, columns, source)
	if err != nil {
		return 0, err
	}
	return copyCount, nil
}
