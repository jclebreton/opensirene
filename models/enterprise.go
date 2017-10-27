package models

// Enterprise is a full struct containing all the possible information about an
// entry in the database
type Enterprise struct {

	// Identification
	Siret int64  `json:"siret"` // Combinaison du SIREN + NIC
	Siren string `json:"siren"` // Identifiant de l'entreprise
	Nic   string `json:"nic"`   // Numéro interne de classement de l’établissement

	// Adressage Normalisé
	L1Normalisee string `json:"l1_normalisee"`
	L2Normalisee string `json:"l2_normalisee"`
	L3Normalisee string `json:"l3_normalisee"`
	L4Normalisee string `json:"l4_normalisee"`
	L5Normalisee string `json:"l5_normalisee"`
	L6Normalisee string `json:"l6_normalisee"`
	L7Normalisee string `json:"l7_normalisee"`

	// Adressage Déclaré
	L1Declaree string `json:"l1_declaree"`
	L2Declaree string `json:"l2_declaree"`
	L3Declaree string `json:"l3_declaree"`
	L4Declaree string `json:"l4_declaree"`
	L5Declaree string `json:"l5_declaree"`
	L6Declaree string `json:"l6_declaree"`
	L7Declaree string `json:"l7_declaree"`

	// Adressage Géographique
	NumVoie string // Numéro dans la voie
	IndRep  string // Indice de répétition
	TypVoie string // Type de la voie
	LibVoie string // Libellé de la voie
	CodePos string // Code postal
	Cedex   string // Code Cedex

	// Localisation Géographique
	RPEt    string // Région de localisation
	LibReg  string // Libellé de la région
	DepEt   string // Département
	ArronEt string // Arrondissement
	CtonEt  string // Canton
	ComEt   string // Commune
	LibCom  string // Libellé de la commune de localisation
	DU      string // Département de l'unité urbaine de la localisation
	TU      string // Taille de l'unité urbaine
	UU      string // Numéro de l'unité urbaine
	EPCI    string // Localisation dans un établissement public de coopération intercommunale
	TCD     string // Tranche de commune détaillée
	ZemEt   string // Zone d'emploi

	// Informations
	Siege      string // Qualité de siège ou non de l'établissement
	Enseigne   string // Enseigne ou nom de l'exploitation
	IndPublipo string // Indicateur du champ du publipostage
	DiffCom    string // Statut de diffusion de l'établissement
	AmintrEt   string // Année et mois d'introduction de l'établissement dans la base de diffusion

	// Caractéristiques Économiques
	NatEtab    string // Nature de l'établissement d'un entrepreneur individuel
	LibNatEtab string // Libellé de la nature de l'établissement
	APET700    string // Activité principale de l'établissement

	// "numvoie", "indrep", "typvoie", "libvoie", "codpos", "cedex"

	// "siret", "siren", "nic", "l1_normalisee", "l2_normalisee", "l3_normalisee", "l4_normalisee",
	// "l5_normalisee", "l6_normalisee", "l7_normalisee", "l1_declaree", "l2_declaree", "l3_declaree", "l4_declaree",
	// "l5_declaree", "l6_declaree", "l7_declaree", "numvoie", "indrep", "typvoie", "libvoie", "codpos", "cedex",

	// "rpet", "libreg", "depet", "arronet", "ctonet", "comet", "libcom", "du", "tu", "uu", "epci", "tcd", "zemet",
	// "siege", "enseigne", "ind_publipo", "diffcom", "amintret", "natetab", "libnatetab", "apet700",

	// "libapet",
	// "dapet", "tefet", "libtefet", "efetcent", "defet", "origine", "dcret", "ddebact", "activnat", "lieuact",
	// "actisurf", "saisonat", "modet", "prodet", "prodpart", "auxilt", "nomen_long", "sigle", "nom", "prenom",
	// "civilite", "rna", "nicsiege", "rpen", "depcomen", "adr_mail", "nj", "libnj", "apen700", "libapen", "dapen",
	// "aprm", "ess", "dateess", "tefen", "libtefen", "efencent", "defen", "categorie", "dcren", "amintren", "monoact",
	// "moden", "proden", "esaann", "tca", "esaapen", "esasec1n", "esasec2n", "esasec3n", "esasec4n", "vmaj", "vmaj1",
	// "vmaj2", "vmaj3", "datemaj"
}
