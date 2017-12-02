package establishments

import (
	"time"

	"github.com/jclebreton/opensirene/domain"
)

type Establishment struct {
	// Identification
	Siren string `gorm:"column:siren"` // Identifiant de l'entreprise
	Nic   string `gorm:"column:nic"`   // Numéro interne de classement de l’établissement

	// Adressage Normalisé
	L1Normalisee string `gorm:"column:l1_normalisee"`
	L2Normalisee string `gorm:"column:l2_normalisee"`
	L3Normalisee string `gorm:"column:l3_normalisee"`
	L4Normalisee string `gorm:"column:l4_normalisee"`
	L5Normalisee string `gorm:"column:l5_normalisee"`
	L6Normalisee string `gorm:"column:l6_normalisee"`
	L7Normalisee string `gorm:"column:l7_normalisee"`

	// Adressage Déclaré
	L1Declaree string `gorm:"column:l1_declaree"`
	L2Declaree string `gorm:"column:l2_declaree"`
	L3Declaree string `gorm:"column:l3_declaree"`
	L4Declaree string `gorm:"column:l4_declaree"`
	L5Declaree string `gorm:"column:l5_declaree"`
	L6Declaree string `gorm:"column:l6_declaree"`
	L7Declaree string `gorm:"column:l7_declaree"`

	// Adressage Géographique
	NumVoie string `gorm:"column:numvoie"` // Numéro dans la voie
	IndRep  string `gorm:"column:indrep"`  // Indice de répétition
	TypVoie string `gorm:"column:typvoie"` // Type de la voie
	LibVoie string `gorm:"column:libvoie"` // Libellé de la voie
	CodePos string `gorm:"column:codepos"` // Code postal
	Cedex   string `gorm:"column:cedex"`   // Code Cedex

	// Localisation Géographique
	RPEt    string `gorm:"column:rpet"`    // Région de localisation
	LibReg  string `gorm:"column:libreg"`  // Libellé de la région
	DepEt   string `gorm:"column:depet"`   // Département
	ArronEt string `gorm:"column:arronet"` // Arrondissement
	CtonEt  string `gorm:"column:ctonet"`  // Canton
	ComEt   string `gorm:"column:comet"`   // Commune
	LibCom  string `gorm:"column:libcom"`  // Libellé de la commune de localisation
	DU      string `gorm:"column:du"`      // Département de l'unité urbaine de la localisation
	TU      string `gorm:"column:tu"`      // Taille de l'unité urbaine
	UU      string `gorm:"column:uu"`      // Numéro de l'unité urbaine
	EPCI    string `gorm:"column:epci"`    // Localisation dans un établissement public de coopération intercommunale
	TCD     string `gorm:"column:tcd"`     // Tranche de commune détaillée
	ZemEt   string `gorm:"column:zemet"`   // Zone d'emploi

	// Informations
	Siege      string     `gorm:"column:siege"`       // Qualité de siège ou non de l'établissement
	Enseigne   string     `gorm:"column:enseigne"`    // Enseigne ou nom de l'exploitation
	IndPublipo string     `gorm:"column:ind_publipo"` // Indicateur du champ du publipostage
	DiffCom    string     `gorm:"column:diffcom"`     // Statut de diffusion de l'établissement
	AmintrEt   *time.Time `gorm:"column:amintret"`    // Année et mois d'introduction de l'établissement dans la base de diffusion

	// Caractéristiques Économiques
	NatEtab    string     `gorm:"column:natetab"`    // Nature de l'établissement d'un entrepreneur individuel
	LibNatEtab string     `gorm:"column:libnatetab"` // Libellé de la nature de l'établissement
	APET700    string     `gorm:"column:apet700"`    // Activité principale de l'établissement
	LibAPET    string     `gorm:"column:libapet"`    // Libellé de l'activité principale de l'établissement
	DAPET      *time.Time `gorm:"column:dapet"`      // Année de validité de l'activité principale de l'établissement
	TEfet      string     `gorm:"column:tefet"`      // Tranche d'effectif salarié de l'établissement
	LibTEfet   string     `gorm:"column:libtefet"`   // Libellé de la tranche d'effectif de l'établissement
	EfetCent   string     `gorm:"column:efetcent"`   // Effectif salarié de l'établissement à la centaine près
	DEfet      *time.Time `gorm:"column:defet"`      // Année de validité de l'effectif salarié de l'établissement
	Origine    string     `gorm:"column:origine"`    // Origine de la création de l'établissement
	DCret      *time.Time `gorm:"column:dcret"`      // Année et mois de création de l'établissement
	DDebAct    *time.Time `gorm:"column:ddebact"`    // Date de début d’activité
	ActivNat   string     `gorm:"column:activnat"`   // Nature de l'activité de l'établissement
	LieuAct    string     `gorm:"column:lieuact"`    // Lieu de l'activité de l'établissement
	ActiSurf   string     `gorm:"column:actisurf"`   // Type de magasin
	SaisonAt   string     `gorm:"column:saisonat"`   // Caractère saisonnier ou non de l'activité de l'établissement
	ModEt      string     `gorm:"column:modet"`      // Modalité de l'activité principale de l'établissement
	ProdEt     string     `gorm:"column:prodet"`     // Caractère productif de l'établissement
	ProdPart   string     `gorm:"column:prodpart"`   // Participation particulière à la production de l'établissement
	AuxiLt     string     `gorm:"column:auxilt"`     // Caractère auxiliaire de l'activité de l'établissement

	// Identification de l'entreprise
	NomenLong string `gorm:"column:nomen_long"` // Nom ou raison sociale de l'entreprise
	Sigle     string `gorm:"column:sigle"`      // Sigle de l'entreprise
	Nom       string `gorm:"column:nom"`        // Nom de naissance
	Prenom    string `gorm:"column:prenom"`     // Prénom
	Civilite  string `gorm:"column:civilite"`   // Civilité des entrepreneurs individuels
	RNA       string `gorm:"column:rna"`        // Numéro d’identification au répertoire national des associations

	// Informations sur le siègle de l'entreprise
	NicSiege string `gorm:"column:nicsiege"` // Numéro interne de classement de l'établissement siège
	RPEn     string `gorm:"column:rpen"`     // Région de localisation du siège de l'entreprise
	DepComEn string `gorm:"column:depcomen"` //Département et commune de localisation du siège de l'entreprise
	AdrMail  string `gorm:"column:adr_mail"` // Adresse mail

	// Caractéristiques économiques de l'entreprise
	NJ        string     `gorm:"column:nj"`      // Nature juridique de l'entreprise
	LibNJ     string     `gorm:"column:libnj"`   // Libellé de la nature juridique
	APEN700   string     `gorm:"column:apen700"` // Activité principale de l'entreprise
	LibAPEN   string     `gorm:"column:libapen"` // Libellé de l'activité principale de l'entreprise
	DAPEN     *time.Time `gorm:"column:dapen"`   // Année de validité de l'activité principale de l'entreprise
	APRM      string     `gorm:"column:aprm"`    // Activité principale au registre des métiers
	ESS       string     `gorm:"column:ess"`     // Appartenance au champ de l’économie sociale et solidaire
	DateESS   *time.Time
	TefEn     string     `gorm:"column:tefen"`     // Tranche d'effectif salarié de l'entreprise
	LibTefEn  string     `gorm:"column:libtefen"`  // Libellé de la tranche d'effectif de l'entreprise
	EfEnCent  string     `gorm:"column:efencent"`  // Effectif salarié de l'entreprise à la centaine près
	DEfEn     *time.Time `gorm:"column:defen"`     // Année de validité de l'effectif salarié de l'entreprise
	Categorie string     `gorm:"column:categorie"` // Catégorie d'entreprise
	DCrEn     *time.Time `gorm:"column:dcren"`     // Année et mois de création de l'entreprise
	AmintrEn  *time.Time `gorm:"column:amintren"`  // Année et mois d'introduction de l'entreprise dans la base de diffusion
	MonoAct   string     `gorm:"column:monoact"`   // Indice de monoactivité de l'entreprise
	ModEn     string     `gorm:"column:moden"`     // Modalité de l'activité principale de l'entreprise
	ProdEn    string     `gorm:"column:proden"`    // Caractère productif de l'entreprise
	ESAANN    *time.Time `gorm:"column:esaann"`    // Année de validité des rubriques de niveau entreprise en provenance de l'ESA*
	TCA       string     `gorm:"column:tca"`       // Tranche de chiffre d'affaires pour les entreprises enquêtées par l'ESA*
	ESAAPEN   string     `gorm:"column:esaapen"`   // Activité principale de l'entreprise issue de l'ESA*
	ESASEC1N  string     `gorm:"column:esasec1n"`  // Première activité secondaire déclarée dans l'ESA*
	ESASEC2N  string     `gorm:"column:esasec2n"`  // Deuxième activité secondaire déclarée dans l'ESA*
	ESASEC3N  string     `gorm:"column:esasec3n"`  // Troisième activité secondaire déclarée dans l'ESA*
	ESASEC4N  string     `gorm:"column:esasec4n"`  // Quatrième activité secondaire déclarée dans l'ESA*

	// Données specifiques aux mises à jour
	VMaj    string     `gorm:"column:vmaj"`    // Nature de la mise à jour (création, suppression, modification)
	VMaj1   string     `gorm:"column:vmaj1"`   // Indicateur de mise à jour n°1
	VMaj2   string     `gorm:"column:vmaj2"`   // Indicateur de mise à jour n°2
	VMaj3   string     `gorm:"column:vmaj3"`   // Indicateur de mise à jour n°3
	DateMaj *time.Time `gorm:"column:datemaj"` // Date de traitement de la mise à jour
}

func (e *Establishment) ToUC() *domain.Establishment {
	return &domain.Establishment{
		Siret:        e.Siren + e.Nic,
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

// Enterprises is a slice of Enterprise
type Establishments []Establishment

func (es *Establishments) ToUC() *[]domain.Establishment {
	var result []domain.Establishment
	for _, e := range *es {
		result = append(result, *e.ToUC())
	}
	return &result
}
