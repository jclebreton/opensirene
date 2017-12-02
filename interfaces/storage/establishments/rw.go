package establishments

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/jackc/pgx"
	"github.com/jclebreton/opensirene/domain"
	"github.com/jinzhu/gorm"
)

type RW struct {
	GormClient *gorm.DB
	PgxClient  *pgx.ConnPool
}

var cols = []string{
	"siren", "nic", "l1_normalisee", "l2_normalisee", "l3_normalisee", "l4_normalisee",
	"l5_normalisee", "l6_normalisee", "l7_normalisee", "l1_declaree", "l2_declaree", "l3_declaree", "l4_declaree",
	"l5_declaree", "l6_declaree", "l7_declaree", "numvoie", "indrep", "typvoie", "libvoie", "codpos", "cedex",
	"rpet", "libreg", "depet", "arronet", "ctonet", "comet", "libcom", "du", "tu", "uu", "epci", "tcd", "zemet",
	"siege", "enseigne", "ind_publipo", "diffcom", "amintret", "natetab", "libnatetab", "apet700", "libapet",
	"dapet", "tefet", "libtefet", "efetcent", "defet", "origine", "dcret", "ddebact", "activnat", "lieuact",
	"actisurf", "saisonat", "modet", "prodet", "prodpart", "auxilt", "nomen_long", "sigle", "nom", "prenom",
	"civilite", "rna", "nicsiege", "rpen", "depcomen", "adr_mail", "nj", "libnj", "apen700", "libapen", "dapen",
	"aprm", "ess", "dateess", "tefen", "libtefen", "efencent", "defen", "categorie", "dcren", "amintren", "monoact",
	"moden", "proden", "esaann", "tca", "esaapen", "esasec1n", "esasec2n", "esasec3n", "esasec4n", "vmaj", "vmaj1",
	"vmaj2", "vmaj3", "datemaj",
}

func (rw *RW) FindEnterpriseBySiret(siret string) (*domain.Establishment, error) {
	e := domain.Establishment{}
	if err := e.SetSiret(siret); err != nil {
		return nil, err
	}

	sql := fmt.Sprintf("SELECT %s FROM enterprises WHERE siren=$1 AND nic=$2", strings.Join(cols, ", "))
	err := rw.PgxClient.QueryRow(sql, e.Siren, e.Nic).Scan(&e.Siren, &e.Nic, &e.L1Normalisee, &e.L2Normalisee,
		&e.L3Normalisee, &e.L4Normalisee, &e.L5Normalisee, &e.L6Normalisee, &e.L7Normalisee, &e.L1Declaree,
		&e.L2Declaree, &e.L3Declaree, &e.L4Declaree, &e.L5Declaree, &e.L6Declaree, &e.L7Declaree, &e.NumVoie, &e.IndRep,
		&e.TypVoie, &e.LibVoie, &e.CodePos, &e.Cedex, &e.RPEt, &e.LibReg, &e.DepEt, &e.ArronEt, &e.CtonEt, &e.ComEt,
		&e.LibCom, &e.DU, &e.TU, &e.UU, &e.EPCI, &e.TCD, &e.ZemEt, &e.Siege, &e.Enseigne, &e.IndPublipo, &e.DiffCom,
		&e.AmintrEt, &e.NatEtab, &e.LibNatEtab, &e.APET700, &e.LibAPET, &e.DAPET, &e.TEfet, &e.LibTEfet, &e.EfetCent,
		&e.DEfet, &e.Origine, &e.DCret, &e.DDebAct, &e.ActivNat, &e.LieuAct, &e.ActiSurf, &e.SaisonAt, &e.ModEt,
		&e.ProdEt, &e.ProdPart, &e.AuxiLt, &e.NomenLong, &e.Sigle, &e.Nom, &e.Prenom, &e.Civilite, &e.RNA, &e.NicSiege,
		&e.RPEn, &e.DepComEn, &e.AdrMail, &e.NJ, &e.LibNJ, &e.APEN700, &e.LibAPEN, &e.DAPEN, &e.APRM, &e.ESS,
		&e.DateESS, &e.TefEn, &e.LibTefEn, &e.EfEnCent, &e.DEfEn, &e.Categorie, &e.DCrEn, &e.AmintrEn, &e.MonoAct,
		&e.ModEn, &e.ProdEn, &e.ESAANN, &e.TCA, &e.ESAAPEN, &e.ESASEC1N, &e.ESASEC2N, &e.ESASEC3N, &e.ESASEC4N, &e.VMaj,
		&e.VMaj1, &e.VMaj2, &e.VMaj3, &e.DateMaj)
	if err != nil {
		return nil, err
	}

	return &e, nil
}

func (rw *RW) FindEstablishmentsFromSiren(siren, limit, offset string) (*[]domain.Establishment, error) {
	var err error
	var es Establishments
	lim := -1
	off := -1

	if limit != "" {
		if lim, err = strconv.Atoi(limit); err != nil {
			return nil, errors.New("'limit' query parameter isn't an integer")
		}
	}
	if offset != "" {
		if off, err = strconv.Atoi(offset); err != nil {
			return nil, errors.New("'offset' query parameter isn't an integer")
		}
	}

	res := rw.GormClient.Limit(lim).Offset(off).Order("nic ASC").Find(&es, Establishment{Siren: siren})
	if res.RecordNotFound() || len(es) == 0 {
		return nil, errors.New("not found")
	}
	if res.Error != nil {
		return nil, res.Error
	}

	return es.ToUC(), nil
}
