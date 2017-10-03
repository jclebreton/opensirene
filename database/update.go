package database

import "github.com/jclebreton/opensirene/download-extract"

type Update struct {
	db *DBClient
}

func InitUpdate(db *DBClient) *Update {
	return &Update{db: db}
}

func (update *Update) ImportCompleteUpdateFile(path string, progress chan download_extract.Progression) (int, error) {
	copyFromSource, err := InitCopyFromFile(path, progress)
	if err != nil {
		return 0, err
	}

	columns := []string{"siren", "nic", "l1_normalisee", "l2_normalisee", "l3_normalisee", "l4_normalisee",
		"l5_normalisee", "l6_normalisee", "l7_normalisee", "l1_declaree", "l2_declaree", "l3_declaree", "l4_declaree",
		"l5_declaree", "l6_declaree", "l7_declaree", "numvoie", "indrep", "typvoie", "libvoie", "codpos", "cedex",
		"rpet", "libreg", "depet", "arronet", "ctonet", "comet", "libcom", "du", "tu", "uu", "epci", "tcd", "zemet",
		"siege", "enseigne", "ind_publipo", "diffcom", "amintret", "natetab", "libnatetab", "apet700", "libapet",
		"dapet", "tefet", "libtefet", "efetcent", "defet", "origine", "dcret", "ddebact", "activnat", "lieuact",
		"actisurf", "saisonat", "modet", "prodet", "prodpart", "auxilt", "nomen_long", "sigle", "nom", "prenom",
		"civilite", "rna", "nicsiege", "rpen", "depcomen", "adr_mail", "nj", "libnj", "apen700", "libapen", "dapen",
		"aprm", "ess", "dateess", "tefen", "libtefen", "efencent", "defen", "categorie", "dcren", "amintren", "monoact",
		"moden", "proden", "esaann", "tca", "esaapen", "esasec1n", "esasec2n", "esasec3n", "esasec4n", "vmaj", "vmaj1",
		"vmaj2", "vmaj3", "datemaj"}

	copyCount, err := update.db.ImportCSVFile("enterprises", columns, copyFromSource)
	if err != nil {
		return 0, err
	}

	return copyCount, nil
}

func (update *Update) ImportIncrementalUpdateFile(path string, progress chan download_extract.Progression) (int, error) {
	copyFromSource, err := InitCopyFromFile(path, progress)
	if err != nil {
		return 0, err
	}

	columns := []string{"siren", "nic", "l1_normalisee", "l2_normalisee", "l3_normalisee", "l4_normalisee",
		"l5_normalisee", "l6_normalisee", "l7_normalisee", "l1_declaree", "l2_declaree", "l3_declaree", "l4_declaree",
		"l5_declaree", "l6_declaree", "l7_declaree", "numvoie", "indrep", "typvoie", "libvoie", "codpos", "cedex",
		"rpet", "libreg", "depet", "arronet", "ctonet", "comet", "libcom", "du", "tu", "uu", "epci", "tcd", "zemet",
		"siege", "enseigne", "ind_publipo", "diffcom", "amintret", "natetab", "libnatetab", "apet700", "libapet",
		"dapet", "tefet", "libtefet", "efetcent", "defet", "origine", "dcret", "ddebact", "activnat", "lieuact",
		"actisurf", "saisonat", "modet", "prodet", "prodpart", "auxilt", "nomen_long", "sigle", "nom", "prenom",
		"civilite", "rna", "nicsiege", "rpen", "depcomen", "adr_mail", "nj", "libnj", "apen700", "libapen", "dapen",
		"aprm", "ess", "dateess", "tefen", "libtefen", "efencent", "defen", "categorie", "dcren", "amintren", "monoact",
		"moden", "proden", "esaann", "tca", "esaapen", "esasec1n", "esasec2n", "esasec3n", "esasec4n", "vmaj", "vmaj1",
		"vmaj2", "vmaj3", "datemaj", "eve", "dateve", "typcreh", "dreactet", "dreacten", "madresse", "menseigne",
		"mapet", "mprodet", "mauxilt", "mnomen", "msigle", "mnicsiege", "mnj", "mapen", "mproden", "siretps", "tel"}

	copyCount, err := update.db.ImportCSVFile("temp_incremental", columns, copyFromSource)
	if err != nil {
		return 0, err
	}

	return copyCount, nil
}
