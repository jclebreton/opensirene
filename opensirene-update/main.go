package main

import (
	"fmt"
	"sync"

	"io/ioutil"
	"os"
)

var url = "http://files.data.gouv.fr/sirene"
var workingDirectory = "/tmp/"

func main() {

	//Working directory
	dest, err := ioutil.TempDir(workingDirectory, "tmp")
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	fmt.Printf("Working directory: %s\n", dest)

	//Files to download
	files := getIncrementalFiles()
	files = append(files, getLastStockFile())

	//Parsing each file
	var wg sync.WaitGroup
	for _, file := range files {
		go func(file string) {

			wg.Add(1)
			defer wg.Done()

			//Download
			fmt.Printf("Downloading %s.zip\n", file)
			err := downloadFile(file+".zip", url, dest)
			if err != nil {
				fmt.Printf("Error: %s", err)
				return
			}

			//Unzip
			fmt.Printf("Unzipping %s.zip\n", file)
			_, err = unzipFile(dest+"/"+file+".zip", dest)
			if err != nil {
				fmt.Printf("Error: %s", err)
				return
			}

			//Remove zip
			fmt.Printf("Removing %s.zip\n", file)
			err = os.Remove(dest + "/" + file + ".zip")
			if err != nil {
				fmt.Printf("Error: %s", err)
				return
			}
		}(file)
	}
	wg.Wait()

	fmt.Println("All CSV files are availables.")
}
