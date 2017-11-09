package json

import (
	"time"

	"github.com/jclebreton/opensirene/domain"
)

type formatGetEnterprisesResp struct{}

type establishmentJSON struct {
	Siret        string    `json:"siret"`
	Siren        string    `json:"siren"`
	Nic          string    `json:"nic"`
	L1Normalisee string    `json:"l1_normalisee"`
	L2Normalisee string    `json:"l2_normalisee"`
	L3Normalisee string    `json:"l3_normalisee"`
	L4Normalisee string    `json:"l4_normalisee"`
	L5Normalisee string    `json:"l5_normalisee"`
	L6Normalisee string    `json:"l6_normalisee"`
	L7Normalisee string    `json:"l7_normalisee"`
	L1Declaree   string    `json:"l1_declaree"`
	L2Declaree   string    `json:"l2_declaree"`
	L3Declaree   string    `json:"l3_declaree"`
	L4Declaree   string    `json:"l4_declaree"`
	L5Declaree   string    `json:"l5_declaree"`
	L6Declaree   string    `json:"l6_declaree"`
	L7Declaree   string    `json:"l7_declaree"`
	NumVoie      string    `json:"numvoie"`
	IndRep       string    `json:"indrep"`
	TypVoie      string    `json:"typvoie"`
	LibVoie      string    `json:"libvoie"`
	CodePos      string    `json:"codepos"`
	Cedex        string    `json:"cedex"`
	RPEt         string    `json:"rpet"`
	LibReg       string    `json:"libreg"`
	DepEt        string    `json:"depet"`
	ArronEt      string    `json:"arronet"`
	CtonEt       string    `json:"ctonet"`
	ComEt        string    `json:"comet"`
	LibCom       string    `json:"libcom"`
	DU           string    `json:"du"`
	TU           string    `json:"tu"`
	UU           string    `json:"uu"`
	EPCI         string    `json:"epci"`
	TCD          string    `json:"tcd"`
	ZemEt        string    `json:"zem_et"`
	Siege        string    `json:"siege"`
	Enseigne     string    `json:"enseigne"`
	IndPublipo   string    `json:"ind_publipo"`
	DiffCom      string    `json:"diffcom"`
	AmintrEt     string    `json:"amintret"`
	NatEtab      string    `json:"natetab"`
	LibNatEtab   string    `json:"libnatetab"`
	APET700      string    `json:"apet700"`
	LibAPET      string    `json:"libapet"`
	DAPET        time.Time `json:"dapet"`
	TEfet        string    `json:"tefet"`
	LibTEfet     string    `json:"libtefet"`
	EfetCent     string    `json:"efetcent"`
	DEfet        time.Time `json:"defet"`
	Origine      string    `json:"origine"`
	DCret        string    `json:"dcret"`
	DDebAct      time.Time `json:"ddebact"`
	ActivNat     string    `json:"activnat"`
	LieuAct      string    `json:"lieu_act"`
	ActiSurf     string    `json:"actisurf"`
	SaisonAt     string    `json:"saisonat"`
	ModEt        string    `json:"modet"`
	ProdEt       string    `json:"prodet"`
	ProdPart     string    `json:"prodpart"`
	AuxiLt       string    `json:"auxilt"`
	NomenLong    string    `json:"nomen_long"`
	Sigle        string    `json:"sigle"`
	Nom          string    `json:"nom"`
	Prenom       string    `json:"prenom"`
	Civilite     string    `json:"civilite"`
	RNA          string    `json:"rna"`
	NicSiege     string    `json:"nicsiege"`
	RPEn         string    `json:"rpen"`
	DepComEn     string    `json:"depcomen"`
	AdrMail      string    `json:"adr_mail"`
	NJ           string    `json:"nj"`
	LibNJ        string    `json:"libnj"`
	APEN700      string    `json:"apen700"`
	LibAPEN      string    `json:"libapen"`
	DAPEN        time.Time `json:"dapen"`
	APRM         string    `json:"aprm"`
	ESS          string    `json:"ess"`
	TefEn        string    `json:"tefen"`
	LibTefEn     string    `json:"libtefen"`
	EfEnCent     string    `json:"efencent"`
	DEfEn        time.Time `json:"defen"`
	Categorie    string    `json:"categorie"`
	DCrEn        string    `json:"dcren"`
	AmintrEn     string    `json:"amintren"`
	MonoAct      string    `json:"monoact"`
	ModEn        string    `json:"moden"`
	ProdEn       string    `json:"proden"`
	ESAANN       time.Time `json:"esaann"`
	TCA          string    `json:"tca"`
	ESAAPEN      string    `json:"esaapen"`
	ESASEC1N     string    `json:"esasec1n"`
	ESASEC2N     string    `json:"esasec2n"`
	ESASEC3N     string    `json:"esasec3n"`
	ESASEC4N     string    `json:"esasec4n"`
	VMaj         string    `json:"vmaj"`
	VMaj1        string    `json:"vmaj1"`
	VMaj2        string    `json:"vmaj2"`
	VMaj3        string    `json:"vmaj3"`
	DateMaj      time.Time `json:"datemaj"`
}

func (jw *formatGetEnterprisesResp) FormatGetEnterpriseFromSiretResp(e domain.Establishment) interface{} {
	return jw.fromUC(e)
}

func (jw *formatGetEnterprisesResp) FormatGetEstablishmentsFromSirenResp(es []domain.Establishment) interface{} {
	var eeJSON []establishmentJSON
	for _, e := range es {
		eeJSON = append(eeJSON, jw.fromUC(e))
	}
	return eeJSON
}

func (jw *formatGetEnterprisesResp) fromUC(e domain.Establishment) establishmentJSON {
	return establishmentJSON{
		Siret:        e.Siret,
		Siren:        e.Siren,
		Nic:          e.Nic,
		L1Normalisee: e.L1Normalisee,
		L2Normalisee: e.L2Normalisee,
		L3Normalisee: e.L3Normalisee,
		L4Normalisee: e.L4Normalisee,
		L5Normalisee: e.L5Normalisee,
		L6Normalisee: e.L6Normalisee,
		L7Normalisee: e.L7Normalisee,
		L1Declaree:   e.L1Declaree,
		L2Declaree:   e.L2Declaree,
		L3Declaree:   e.L3Declaree,
		L4Declaree:   e.L4Declaree,
		L5Declaree:   e.L5Declaree,
		L6Declaree:   e.L6Declaree,
		L7Declaree:   e.L7Declaree,
		NumVoie:      e.NumVoie,
		IndRep:       e.IndRep,
		TypVoie:      e.TypVoie,
		LibVoie:      e.LibVoie,
		CodePos:      e.CodePos,
		Cedex:        e.Cedex,
		RPEt:         e.RPEt,
		LibReg:       e.LibReg,
		DepEt:        e.DepEt,
		ArronEt:      e.ArronEt,
		CtonEt:       e.CtonEt,
		ComEt:        e.ComEt,
		LibCom:       e.LibCom,
		DU:           e.DU,
		TU:           e.TU,
		UU:           e.UU,
		EPCI:         e.EPCI,
		TCD:          e.TCD,
		ZemEt:        e.ZemEt,
		Siege:        e.Siege,
		Enseigne:     e.Enseigne,
		IndPublipo:   e.IndPublipo,
		DiffCom:      e.DiffCom,
		AmintrEt:     e.AmintrEt,
		NatEtab:      e.NatEtab,
		LibNatEtab:   e.LibNatEtab,
		APET700:      e.APET700,
		LibAPET:      e.LibAPET,
		DAPET:        e.DAPET,
		TEfet:        e.TEfet,
		LibTEfet:     e.LibTEfet,
		EfetCent:     e.EfetCent,
		DEfet:        e.DEfet,
		Origine:      e.Origine,
		DCret:        e.DCret,
		DDebAct:      e.DDebAct,
		ActivNat:     e.ActivNat,
		LieuAct:      e.LieuAct,
		ActiSurf:     e.ActiSurf,
		SaisonAt:     e.SaisonAt,
		ModEt:        e.ModEt,
		ProdEt:       e.ProdEt,
		ProdPart:     e.ProdPart,
		AuxiLt:       e.AuxiLt,
		NomenLong:    e.NomenLong,
		Sigle:        e.Sigle,
		Nom:          e.Nom,
		Prenom:       e.Prenom,
		Civilite:     e.Civilite,
		RNA:          e.RNA,
		NicSiege:     e.NicSiege,
		RPEn:         e.RPEn,
		DepComEn:     e.DepComEn,
		AdrMail:      e.AdrMail,
		NJ:           e.NJ,
		LibNJ:        e.LibNJ,
		APEN700:      e.APEN700,
		LibAPEN:      e.LibAPEN,
		DAPEN:        e.DAPEN,
		APRM:         e.APRM,
		ESS:          e.ESS,
		TefEn:        e.TefEn,
		LibTefEn:     e.LibTefEn,
		EfEnCent:     e.EfEnCent,
		DEfEn:        e.DEfEn,
		Categorie:    e.Categorie,
		DCrEn:        e.DCrEn,
		AmintrEn:     e.AmintrEn,
		MonoAct:      e.MonoAct,
		ModEn:        e.ModEn,
		ProdEn:       e.ProdEn,
		ESAANN:       e.ESAANN,
		TCA:          e.TCA,
		ESAAPEN:      e.ESAAPEN,
		ESASEC1N:     e.ESASEC1N,
		ESASEC2N:     e.ESASEC2N,
		ESASEC3N:     e.ESASEC3N,
		ESASEC4N:     e.ESASEC4N,
		VMaj:         e.VMaj,
		VMaj1:        e.VMaj1,
		VMaj2:        e.VMaj2,
		VMaj3:        e.VMaj3,
		DateMaj:      e.DateMaj,
	}
}
