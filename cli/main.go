package main

import (
	"os"

	"fmt"

	"github.com/jclebreton/opensirene/download-extract"
	log "github.com/sirupsen/logrus"
)

var url = "http://files.data.gouv.fr/sirene"
var logFile = "opensirene.log"
var nbWorkersMax = 31

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY, 0666)
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
  --maxworkers=<int> Maximum number of workers to use for processing files (min: 1, max: 31)
  --month=<string>   Month to download (ex: Sep)
  --debug            Enable debugging
  -h --help          Show this screen.`

	args := InitArgs(usage)

	if args.isDebug() {
		log.SetLevel(log.DebugLevel)
	}

	//Working directory
	wd, err := args.getWorkingDirectory()
	if err != nil {
		log.Fatal(err)
		return
	}

	//Max workers
	nbWorkers, err := args.getNbWorkers()
	if err != nil {
		log.Fatal(err)
		return
	}

	log.WithFields(log.Fields{
		"Working directory": wd,
		"Number of workers": nbWorkers,
	}).Debug()

	//Update from scratch
	var zipFiles []download_extract.ZipFile
	if args.isCompleteUpdate() {
		zipFiles, err = download_extract.GetScratchZipList(args.getMonth(), url, wd)
	} else if args.isDailyUpdate() {
		zipFiles, err = download_extract.GetDailyZipList(url, wd)
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
		"Filenames":       download_extract.GetZipFileNames(zipFiles),
	}).Info("Zip files to dowload")

	if len(zipFiles) == 0 {
		fmt.Println("No files to download")
		return
	}

	csvFiles, err := download_extract.DownloadAndExtract(zipFiles, nbWorkers)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.WithFields(log.Fields{
		"Number of files": len(csvFiles),
		"Filenames":       download_extract.GetCsvFileNames(csvFiles),
	}).Info("CSV files extracted")

	//db, err := database.InitDBClient()
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
	//
	//copyFromSource := database.InitCopyFromSource("/home/jc/sept/sirc-17804_9075_14209_201708_L_M_20170901_025839232.csv")
	//err = db.ImportStockFile(copyFromSource)
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}

	fmt.Println()
}
