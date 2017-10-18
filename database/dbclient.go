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

//Optimize table: create visibility map
func (client *DBClient) Optimize(table string) error {
	_, err := client.conn.Exec(fmt.Sprintf("VACUUM ANALYZE %s", table))
	return err
}

func (client *DBClient) ApplyIncremental() error {

	//Delete removed (E), exited (O) and modified (I) companies
	_, err := client.conn.Exec(`
		DELETE FROM enterprises e
		USING temp_incremental i
		WHERE (e.siren, e.nic) = (i.siren, i.nic)
		AND i.vmaj IN ('E', 'O', 'I')
	`)
	if err != nil {
		return err
	}

	//Adds new companies (C), which enter (D) and have been modified (F)
	_, err = client.conn.Exec(`
		INSERT INTO enterprises (
			siren, nic, l1_normalisee, l2_normalisee, l3_normalisee, l4_normalisee, l5_normalisee,
			l6_normalisee, l7_normalisee, l1_declaree, l2_declaree, l3_declaree, l4_declaree,
			l5_declaree, l6_declaree, l7_declaree, numvoie, indrep, typvoie, libvoie, codpos,
			cedex, rpet, libreg, depet, arronet, ctonet, comet, libcom, du, tu, uu, epci,
			tcd, zemet, siege, enseigne, ind_publipo, diffcom, amintret, natetab, libnatetab,
			apet700, libapet, dapet, tefet, libtefet, efetcent, defet, origine, dcret, ddebact,
			activnat, lieuact, actisurf, saisonat, modet, prodet, prodpart, auxilt, nomen_long,
			sigle, nom, prenom, civilite, rna, nicsiege, rpen, depcomen, adr_mail, nj, libnj,
			apen700, libapen, dapen, aprm, ess, dateess, tefen, libtefen, efencent, defen,
			categorie, dcren, amintren, monoact, moden, proden, esaann, tca, esaapen, esasec1n,
			esasec2n, esasec3n, esasec4n, vmaj, vmaj1, vmaj2, vmaj3, datemaj
		)
		SELECT i.siren, i.nic, i.l1_normalisee, i.l2_normalisee, i.l3_normalisee, i.l4_normalisee, i.l5_normalisee,
			i.l6_normalisee, i.l7_normalisee, i.l1_declaree, i.l2_declaree, i.l3_declaree, i.l4_declaree,
			i.l5_declaree, i.l6_declaree, i.l7_declaree, i.numvoie, i.indrep, i.typvoie, i.libvoie, i.codpos,
			i.cedex, i.rpet, i.libreg, i.depet, i.arronet, i.ctonet, i.comet, i.libcom, i.du, i.tu, i.uu, i.epci,
			i.tcd, i.zemet, i.siege, i.enseigne, i.ind_publipo, i.diffcom, i.amintret, i.natetab, i.libnatetab,
			i.apet700, i.libapet, i.dapet, i.tefet, i.libtefet, i.efetcent, i.defet, i.origine, i.dcret, i.ddebact,
			i.activnat, i.lieuact, i.actisurf, i.saisonat, i.modet, i.prodet, i.prodpart, i.auxilt, i.nomen_long,
			i.sigle, i.nom, i.prenom, i.civilite, i.rna, i.nicsiege, i.rpen, i.depcomen, i.adr_mail, i.nj, i.libnj,
			i.apen700, i.libapen, i.dapen, i.aprm, i.ess, i.dateess, i.tefen, i.libtefen, i.efencent, i.defen,
			i.categorie, i.dcren, i.amintren, i.monoact, i.moden, i.proden, i.esaann, i.tca, i.esaapen, i.esasec1n,
			i.esasec2n, i.esasec3n, i.esasec4n, i.vmaj, i.vmaj1, i.vmaj2, i.vmaj3, i.datemaj
		FROM temp_incremental i
		WHERE i.vmaj in ('C', 'D', 'F')
	`)
	if err != nil {
		return err
	}

	return nil
}
