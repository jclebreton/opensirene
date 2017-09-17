package main

import (
	"fmt"

	"io/ioutil"
)

var url = "http://files.data.gouv.fr/sirene"
var wd = "/tmp/"
var nbWorkers = 10

func main() {
	dest, err := ioutil.TempDir(wd, "tmp")
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	csvFiles := downloadAndExtract(dest)
	fmt.Printf("\nNumber of CSV files: %d\n", len(csvFiles))
	fmt.Printf("Process completed\n")
}
