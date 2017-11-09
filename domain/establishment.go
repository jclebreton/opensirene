package domain

import (
	"time"

	"github.com/pkg/errors"
)

// Establishment is a full struct containing all the establishment data
type Establishment struct {
	// Identification
	Siret string // Combinaison du SIREN + NIC
	Siren string // Identifiant de l'entreprise
	Nic   string // Numéro interne de classement de l’établissement

	// Adressage Normalisé
	L1Normalisee string
	L2Normalisee string
	L3Normalisee string
	L4Normalisee string
	L5Normalisee string
	L6Normalisee string
	L7Normalisee string

	// Adressage Déclaré
	L1Declaree string
	L2Declaree string
	L3Declaree string
	L4Declaree string
	L5Declaree string
	L6Declaree string
	L7Declaree string

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
	NatEtab    string    // Nature de l'établissement d'un entrepreneur individuel
	LibNatEtab string    // Libellé de la nature de l'établissement
	APET700    string    // Activité principale de l'établissement
	LibAPET    string    // Libellé de l'activité principale de l'établissement
	DAPET      time.Time // Année de validité de l'activité principale de l'établissement
	TEfet      string    // Tranche d'effectif salarié de l'établissement
	LibTEfet   string    // Libellé de la tranche d'effectif de l'établissement
	EfetCent   string    // Effectif salarié de l'établissement à la centaine près
	DEfet      time.Time // Année de validité de l'effectif salarié de l'établissement
	Origine    string    // Origine de la création de l'établissement
	DCret      string    // Année et mois de création de l'établissement
	DDebAct    time.Time // Date de début d’activité
	ActivNat   string    // Nature de l'activité de l'établissement
	LieuAct    string    // Lieu de l'activité de l'établissement
	ActiSurf   string    // Type de magasin
	SaisonAt   string    // Caractère saisonnier ou non de l'activité de l'établissement
	ModEt      string    // Modalité de l'activité principale de l'établissement
	ProdEt     string    // Caractère productif de l'établissement
	ProdPart   string    // Participation particulière à la production de l'établissement
	AuxiLt     string    // Caractère auxiliaire de l'activité de l'établissement

	// Identification de l'entreprise
	NomenLong string // Nom ou raison sociale de l'entreprise
	Sigle     string // Sigle de l'entreprise
	Nom       string // Nom de naissance
	Prenom    string // Prénom
	Civilite  string // Civilité des entrepreneurs individuels
	RNA       string // Numéro d’identification au répertoire national des associations

	// Informations sur le siègle de l'entreprise
	NicSiege string // Numéro interne de classement de l'établissement siège
	RPEn     string // Région de localisation du siège de l'entreprise
	DepComEn string //Département et commune de localisation du siège de l'entreprise
	AdrMail  string // Adresse mail

	// Caractéristiques économiques de l'entreprise
	NJ        string    // Nature juridique de l'entreprise
	LibNJ     string    // Libellé de la nature juridique
	APEN700   string    // Activité principale de l'entreprise
	LibAPEN   string    // Libellé de l'activité principale de l'entreprise
	DAPEN     time.Time // Année de validité de l'activité principale de l'entreprise
	APRM      string    // Activité principale au registre des métiers
	ESS       string    // Appartenance au champ de l’économie sociale et solidaire
	DateESS   time.Time
	TefEn     string    // Tranche d'effectif salarié de l'entreprise
	LibTefEn  string    // Libellé de la tranche d'effectif de l'entreprise
	EfEnCent  string    // Effectif salarié de l'entreprise à la centaine près
	DEfEn     time.Time // Année de validité de l'effectif salarié de l'entreprise
	Categorie string    // Catégorie d'entreprise
	DCrEn     string    // Année et mois de création de l'entreprise
	AmintrEn  string    // Année et mois d'introduction de l'entreprise dans la base de diffusion
	MonoAct   string    // Indice de monoactivité de l'entreprise
	ModEn     string    // Modalité de l'activité principale de l'entreprise
	ProdEn    string    // Caractère productif de l'entreprise
	ESAANN    time.Time // Année de validité des rubriques de niveau entreprise en provenance de l'ESA*
	TCA       string    // Tranche de chiffre d'affaires pour les entreprises enquêtées par l'ESA*
	ESAAPEN   string    // Activité principale de l'entreprise issue de l'ESA*
	ESASEC1N  string    // Première activité secondaire déclarée dans l'ESA*
	ESASEC2N  string    // Deuxième activité secondaire déclarée dans l'ESA*
	ESASEC3N  string    // Troisième activité secondaire déclarée dans l'ESA*
	ESASEC4N  string    // Quatrième activité secondaire déclarée dans l'ESA*

	// Données specifiques aux mises à jour
	VMaj    string    // Nature de la mise à jour (création, suppression, modification)
	VMaj1   string    // Indicateur de mise à jour n°1
	VMaj2   string    // Indicateur de mise à jour n°2
	VMaj3   string    // Indicateur de mise à jour n°3
	DateMaj time.Time // Date de traitement de la mise à jour
}

// SetSiret is the setter of the Siret and the Nic
func (e *Establishment) SetSiret(siret string) error {
	if len(siret) != 14 {
		return errors.New("not a valid siret")
	}
	e.Siret = siret
	e.Siren = siret[0:9]
	e.Nic = siret[9:14]
	return nil
}
