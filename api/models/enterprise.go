package models

import "time"

// Enterprise is a full struct containing all the possible information about an
// entry in the database
type Enterprise struct {
	// Identification
	Siret string `json:"siret"`                     // Combinaison du SIREN + NIC
	Siren string `json:"siren" gorm:"column:siren"` // Identifiant de l'entreprise
	Nic   string `json:"nic" gorm:"column:nic"`     // Numéro interne de classement de l’établissement

	// Adressage Normalisé
	L1Normalisee string `json:"l1_normalisee" gorm:"column:l1_normalisee"`
	L2Normalisee string `json:"l2_normalisee" gorm:"column:l2_normalisee"`
	L3Normalisee string `json:"l3_normalisee" gorm:"column:l3_normalisee"`
	L4Normalisee string `json:"l4_normalisee" gorm:"column:l4_normalisee"`
	L5Normalisee string `json:"l5_normalisee" gorm:"column:l5_normalisee"`
	L6Normalisee string `json:"l6_normalisee" gorm:"column:l6_normalisee"`
	L7Normalisee string `json:"l7_normalisee" gorm:"column:l7_normalisee"`

	// Adressage Déclaré
	L1Declaree string `json:"l1_declaree" gorm:"column:l1_declaree"`
	L2Declaree string `json:"l2_declaree" gorm:"column:l2_declaree"`
	L3Declaree string `json:"l3_declaree" gorm:"column:l3_declaree"`
	L4Declaree string `json:"l4_declaree" gorm:"column:l4_declaree"`
	L5Declaree string `json:"l5_declaree" gorm:"column:l5_declaree"`
	L6Declaree string `json:"l6_declaree" gorm:"column:l6_declaree"`
	L7Declaree string `json:"l7_declaree" gorm:"column:l7_declaree"`

	// Adressage Géographique
	NumVoie string `json:"numvoie" gorm:"column:numvoie"` // Numéro dans la voie
	IndRep  string `json:"indrep" gorm:"column:indrep"`   // Indice de répétition
	TypVoie string `json:"typvoie" gorm:"column:typvoie"` // Type de la voie
	LibVoie string `json:"libvoie" gorm:"column:libvoie"` // Libellé de la voie
	CodPos  string `json:"codpos" gorm:"column:codpos"`   // Code postal
	Cedex   string `json:"cedex" gorm:"column:cedex"`     // Code Cedex

	// Localisation Géographique
	RPEt    string `json:"rpet" gorm:"column:rpet"`       // Région de localisation
	LibReg  string `json:"libreg" gorm:"column:libreg"`   // Libellé de la région
	DepEt   string `json:"depet" gorm:"column:depet"`     // Département
	ArronEt string `json:"arronet" gorm:"column:arronet"` // Arrondissement
	CtonEt  string `json:"ctonet" gorm:"column:ctonet"`   // Canton
	ComEt   string `json:"comet" gorm:"column:comet"`     // Commune
	LibCom  string `json:"libcom" gorm:"column:libcom"`   // Libellé de la commune de localisation
	DU      string `json:"du" gorm:"column:du"`           // Département de l'unité urbaine de la localisation
	TU      string `json:"tu" gorm:"column:tu"`           // Taille de l'unité urbaine
	UU      string `json:"uu" gorm:"column:uu"`           // Numéro de l'unité urbaine
	EPCI    string `json:"epci" gorm:"column:epci"`       // Localisation dans un établissement public de coopération intercommunale
	TCD     string `json:"tcd" gorm:"column:tcd"`         // Tranche de commune détaillée
	ZemEt   string `json:"zem_et" gorm:"column:zemet"`    // Zone d'emploi

	// Informations
	Siege      string `json:"siege" gorm:"column:siege"`             // Qualité de siège ou non de l'établissement
	Enseigne   string `json:"enseigne" gorm:"column:enseigne"`       // Enseigne ou nom de l'exploitation
	IndPublipo string `json:"ind_publipo" gorm:"column:ind_publipo"` // Indicateur du champ du publipostage
	DiffCom    string `json:"diffcom" gorm:"column:diffcom"`         // Statut de diffusion de l'établissement
	AmintrEt   string `json:"amintret" gorm:"column:amintret"`       // Année et mois d'introduction de l'établissement dans la base de diffusion

	// Caractéristiques Économiques
	NatEtab    string    `json:"natetab" gorm:"column:natetab"`       // Nature de l'établissement d'un entrepreneur individuel
	LibNatEtab string    `json:"libnatetab" gorm:"column:libnatetab"` // Libellé de la nature de l'établissement
	APET700    string    `json:"apet700" gorm:"column:apet700"`       // Activité principale de l'établissement
	LibAPET    string    `json:"libapet" gorm:"column:libapet"`       // Libellé de l'activité principale de l'établissement
	DAPET      time.Time `json:"dapet" gorm:"column:dapet"`           // Année de validité de l'activité principale de l'établissement
	TEfet      string    `json:"tefet" gorm:"column:tefet"`           // Tranche d'effectif salarié de l'établissement
	LibTEfet   string    `json:"libtefet" gorm:"column:libtefet"`     // Libellé de la tranche d'effectif de l'établissement
	EfetCent   string    `json:"efetcent" gorm:"column:efetcent"`     // Effectif salarié de l'établissement à la centaine près
	DEfet      time.Time `json:"defet" gorm:"column:defet"`           // Année de validité de l'effectif salarié de l'établissement
	Origine    string    `json:"origine" gorm:"column:origine"`       // Origine de la création de l'établissement
	DCret      string    `json:"dcret" gorm:"column:dcret"`           // Année et mois de création de l'établissement
	DDebAct    time.Time `json:"ddebact" gorm:"column:ddebact"`       // Date de début d’activité
	ActivNat   string    `json:"activnat" gorm:"column:activnat"`     // Nature de l'activité de l'établissement
	LieuAct    string    `json:"lieu_act" gorm:"column:lieuact"`      // Lieu de l'activité de l'établissement
	ActiSurf   string    `json:"actisurf" gorm:"column:actisurf"`     // Type de magasin
	SaisonAt   string    `json:"saisonat" gorm:"column:saisonat"`     // Caractère saisonnier ou non de l'activité de l'établissement
	ModEt      string    `json:"modet" gorm:"column:modet"`           // Modalité de l'activité principale de l'établissement
	ProdEt     string    `json:"prodet" gorm:"column:prodet"`         // Caractère productif de l'établissement
	ProdPart   string    `json:"prodpart" gorm:"column:prodpart"`     // Participation particulière à la production de l'établissement
	AuxiLt     string    `json:"auxilt" gorm:"column:auxilt"`         // Caractère auxiliaire de l'activité de l'établissement

	// Identification de l'entreprise
	NomenLong string `json:"nomen_long" gorm:"column:nomen_long"` // Nom ou raison sociale de l'entreprise
	Sigle     string `json:"sigle" gorm:"column:sigle"`           // Sigle de l'entreprise
	Nom       string `json:"nom" gorm:"column:nom"`               // Nom de naissance
	Prenom    string `json:"prenom" gorm:"column:prenom"`         // Prénom
	Civilite  string `json:"civilite" gorm:"column:civilite"`     // Civilité des entrepreneurs individuels
	RNA       string `json:"rna" gorm:"column:rna"`               // Numéro d’identification au répertoire national des associations

	// Informations sur le siègle de l'entreprise
	NicSiege string `json:"nicsiege" gorm:"column:nicsiege"` // Numéro interne de classement de l'établissement siège
	RPEn     string `json:"rpen" gorm:"column:rpen"`         // Région de localisation du siège de l'entreprise
	DepComEn string `json:"depcomen" gorm:"column:depcomen"` //Département et commune de localisation du siège de l'entreprise
	AdrMail  string `json:"adr_mail" gorm:"column:adr_mail"` // Adresse mail

	// Caractéristiques économiques de l'entreprise
	NJ        string    `json:"nj" gorm:"column:nj"`               // Nature juridique de l'entreprise
	LibNJ     string    `json:"libnj" gorm:"column:libnj"`         // Libellé de la nature juridique
	APEN700   string    `json:"apen700" gorm:"column:apen700"`     // Activité principale de l'entreprise
	LibAPEN   string    `json:"libapen" gorm:"column:libapen"`     // Libellé de l'activité principale de l'entreprise
	DAPEN     time.Time `json:"dapen" gorm:"column:dapen"`         // Année de validité de l'activité principale de l'entreprise
	APRM      string    `json:"aprm" gorm:"column:aprm"`           // Activité principale au registre des métiers
	ESS       string    `json:"ess" gorm:"column:ess"`             // Appartenance au champ de l’économie sociale et solidaire
	TefEn     string    `json:"tefen" gorm:"column:tefen"`         // Tranche d'effectif salarié de l'entreprise
	LibTefEn  string    `json:"libtefen" gorm:"column:libtefen"`   // Libellé de la tranche d'effectif de l'entreprise
	EfEnCent  string    `json:"efencent" gorm:"column:efencent"`   // Effectif salarié de l'entreprise à la centaine près
	DEfEn     time.Time `json:"defen" gorm:"column:defen"`         // Année de validité de l'effectif salarié de l'entreprise
	Categorie string    `json:"categorie" gorm:"column:categorie"` // Catégorie d'entreprise
	DCrEn     string    `json:"dcren" gorm:"column:dcren"`         // Année et mois de création de l'entreprise
	AmintrEn  string    `json:"amintren" gorm:"column:amintren"`   // Année et mois d'introduction de l'entreprise dans la base de diffusion
	MonoAct   string    `json:"monoact" gorm:"column:monoact"`     // Indice de monoactivité de l'entreprise
	ModEn     string    `json:"moden" gorm:"column:moden"`         // Modalité de l'activité principale de l'entreprise
	ProdEn    string    `json:"proden" gorm:"column:proden"`       // Caractère productif de l'entreprise
	ESAANN    time.Time `json:"esaann" gorm:"column:esaann"`       // Année de validité des rubriques de niveau entreprise en provenance de l'ESA*
	TCA       string    `json:"tca" gorm:"column:tca"`             // Tranche de chiffre d'affaires pour les entreprises enquêtées par l'ESA*
	ESAAPEN   string    `json:"esaapen" gorm:"column:esaapen"`     // Activité principale de l'entreprise issue de l'ESA*
	ESASEC1N  string    `json:"esasec1n" gorm:"column:esasec1n"`   // Première activité secondaire déclarée dans l'ESA*
	ESASEC2N  string    `json:"esasec2n" gorm:"column:esasec2n"`   // Deuxième activité secondaire déclarée dans l'ESA*
	ESASEC3N  string    `json:"esasec3n" gorm:"column:esasec3n"`   // Troisième activité secondaire déclarée dans l'ESA*
	ESASEC4N  string    `json:"esasec4n" gorm:"column:esasec4n"`   // Quatrième activité secondaire déclarée dans l'ESA*

	// Données specifiques aux mises à jour
	VMaj    string    `json:"vmaj" gorm:"column:vmaj"`       // Nature de la mise à jour (création, suppression, modification)
	VMaj1   string    `json:"vmaj1" gorm:"column:vmaj1"`     // Indicateur de mise à jour n°1
	VMaj2   string    `json:"vmaj2" gorm:"column:vmaj2"`     // Indicateur de mise à jour n°2
	VMaj3   string    `json:"vmaj3" gorm:"column:vmaj3"`     // Indicateur de mise à jour n°3
	DateMaj time.Time `json:"datemaj" gorm:"column:datemaj"` // Date de traitement de la mise à jour
}

// Enterprises is a slice of Enterprise
type Enterprises []Enterprise
