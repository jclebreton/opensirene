package main

import (
	"fmt"

	"log"

	"github.com/docopt/docopt-go"
)

var url = "http://files.data.gouv.fr/sirene"
var nbWorkersMax = 31

func main() {
	usage := `Opensirene

French company database based on French government open data.
Github: https://github.com/jclebreton/opensirene

Usage:
  update daily [--wd=<path>]
  update complete [--wd=<path>] [--maxworkers=<int>] [--month=<string>]
  update -h | --help

Options:
  --wd=<path>        Working directory path (by default: /tmp/tmp[0-9]{8,})
  --maxworkers=<int> Maximum number of workers to use for processing files (min: 1, max: 100)
  --month=<string>   Month to download (ex: Sep)
  -h --help          Show this screen.`

	arguments, _ := docopt.Parse(usage, nil, true, "", false)
	//fmt.Println(arguments)

	//Working directory
	wd, err := getWorkingDirectory(arguments)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("Working directory: %s\n", wd)

	//Max workers
	nbWorkers, err := getNbWorkers(arguments)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("Number of workers: %d\n", nbWorkers)

	//Update from scratch
	var zipFiles []zipFile
	if arguments["complete"].(bool) {
		zipFiles, err = getScratchZipList(getMonth(arguments), url, wd)
		fmt.Printf("Prepare complete update\n")
	} else {
		zipFiles, err = getDailyZipList(url, wd)
		fmt.Printf("Prepare daily update\n")
	}

	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("Number of ZIP files: %d\n", len(zipFiles))

	csvFiles, err := downloadAndExtract(zipFiles, nbWorkers)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("\nNumber of CSV files: %d\n", len(csvFiles))

}
