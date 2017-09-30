package main

import (
	"os"

	"github.com/docopt/docopt-go"
	log "github.com/sirupsen/logrus"
)

var url = "http://files.data.gouv.fr/sirene"
var nbWorkersMax = 31

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	file, err := os.OpenFile("opensirene.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		log.SetOutput(file)
	}
	log.SetLevel(log.InfoLevel)
}

func main() {
	usage := `Opensirene

French company database based on French government open data.
Github: https://github.com/jclebreton/opensirene

Usage:
  update daily [--wd=<path>] [--debug]
  update complete [--wd=<path>] [--maxworkers=<int>] [--month=<string>] [--debug]
  update -h | --help

Options:
  --wd=<path>        Working directory path (by default: /tmp/tmp[0-9]{8,})
  --maxworkers=<int> Maximum number of workers to use for processing files (min: 1, max: 100)
  --month=<string>   Month to download (ex: Sep)
  --debug            Enable debugging
  -h --help          Show this screen.`

	arguments, _ := docopt.Parse(usage, nil, true, "", false)

	if arguments["--debug"].(bool) {
		log.SetLevel(log.DebugLevel)
	}

	//Working directory
	wd, err := getWorkingDirectory(arguments)
	if err != nil {
		log.Fatal(err)
		return
	}

	//Max workers
	nbWorkers, err := getNbWorkers(arguments)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.WithFields(log.Fields{
		"Working directory": wd,
		"Number of workers": nbWorkers,
	}).Debug()

	//Update from scratch
	var zipFiles []zipFile
	if arguments["complete"].(bool) {
		zipFiles, err = getScratchZipList(getMonth(arguments), url, wd)
	} else if arguments["daily"].(bool) {
		zipFiles, err = getDailyZipList(url, wd)
	} else {
		log.Fatal("No command selected")
		return
	}

	if err != nil {
		log.Fatal(err)
		return
	}

	log.WithFields(log.Fields{
		"Number of files": len(zipFiles),
		"Filenames":       getZipFileNames(zipFiles),
	}).Info("Zip files to dowload")

	csvFiles, err := downloadAndExtract(zipFiles, nbWorkers)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.WithFields(log.Fields{
		"Number of files": len(csvFiles),
		"Filenames":       getCsvFileNames(csvFiles),
	}).Info("CSV files extracted")
}
