package database

import (
	"fmt"

	"github.com/jackc/pgx"
)

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
			_, err := conn.Exec("SET CLIENT_ENCODING TO 'WIN1252'")
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

func (client *DBClient) ImportStockFile(source pgx.CopyFromSource) error {
	copyCount, err := client.conn.CopyFrom(
		pgx.Identifier{"enterprises"},
		[]string{"siren", "nic", "l1_normalisee", "l2_normalisee", "l3_normalisee", "l4_normalisee", "l5_normalisee",
			"l6_normalisee", "l7_normalisee", "l1_declaree", "l2_declaree", "l3_declaree", "l4_declaree", "l5_declaree",
			"l6_declaree", "l7_declaree", "numvoie", "indrep", "typvoie", "libvoie", "codpos", "cedex", "rpet", "libreg",
			"depet", "arronet", "ctonet", "comet", "libcom", "du", "tu", "uu", "epci", "tcd", "zemet", "siege", "enseigne",
			"ind_publipo", "diffcom", "amintret", "natetab", "libnatetab", "apet700", "libapet", "dapet", "tefet",
			"libtefet", "efetcent", "defet", "origine", "dcret", "ddebact", "activnat", "lieuact", "actisurf", "saisonat",
			"modet", "prodet", "prodpart", "auxilt", "nomen_long", "sigle", "nom", "prenom", "civilite", "rna", "nicsiege",
			"rpen", "depcomen", "adr_mail", "nj", "libnj", "apen700", "libapen", "dapen", "aprm", "ess", "dateess", "tefen",
			"libtefen", "efencent", "defen", "categorie", "dcren", "amintren", "monoact", "moden", "proden", "esaann",
			"tca", "esaapen", "esasec1n", "esasec2n", "esasec3n", "esasec4n", "vmaj", "vmaj1", "vmaj2", "vmaj3", "datemaj"},
		source,
	)

	if err != nil {
		return err
	}

	fmt.Printf("\nTotal: %d", copyCount)

	return nil
}
