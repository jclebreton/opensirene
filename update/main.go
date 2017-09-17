package main

import (
	"fmt"

	"io/ioutil"

	"strconv"

	"github.com/docopt/docopt-go"
)

var url = "http://files.data.gouv.fr/sirene"
var nbWorkers = 100

func main() {
	usage := `Update database from scratch.

Usage:
  update complete [--wd=<path>] [--maxworkers=<int>]
  update -h | --help

Options:
  --wd=<path>        Working directory path (by default: /tmp/tmp[0-9]{8,})
  --maxworkers=<int> Maximum number of workers to use for processing files (min: 1, max: 100)
  -h --help          Show this screen.`

	arguments, _ := docopt.Parse(usage, nil, true, "", false)
	//fmt.Println(arguments)

	//Working directory
	var wd string
	if arguments["--wd"] != nil {
		wd = arguments["--wd"].(string)
	} else {
		temp, err := ioutil.TempDir("/tmp/", "tmp")
		if err != nil {
			fmt.Printf("Error: %s", err)
			return
		}
		wd = temp
	}
	fmt.Printf("Working directory: %s\n", wd)

	//Max workers
	if arguments["--maxworkers"] != nil {
		n, err := strconv.Atoi(arguments["--maxworkers"].(string))
		if err != nil || n <= 0 || n > 100 {
			fmt.Println("Number of workers must be >= 1 and <=100")
			return
		}
		nbWorkers = n
	}
	fmt.Printf("Number of workers: %d\n", nbWorkers)

	//Update from scratch
	if arguments["complete"].(bool) {
		csvFiles := downloadAndExtract(wd)
		fmt.Printf("\nNumber of CSV files: %d\n", len(csvFiles))
		fmt.Printf("Process completed\n")
	}
}
