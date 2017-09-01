package main

import (
	"fmt"
	"os"

	"io/ioutil"

	"sync"
)

var url = "http://files.data.gouv.fr/sirene"
var wd = "/tmp/"

func main() {
	//Working directory
	dest, err := ioutil.TempDir(wd, "tmp")
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	//Get files list
	zipFiles := getZipList(url, dest)

	//Processing files
	fmt.Printf("Processing %d file(s)...\n", len(zipFiles))
	downloadProgress := make(chan map[string]float64)
	unzipProgress := make(chan map[string]float64)
	var wg sync.WaitGroup

	for i := 0; i < len(zipFiles); i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			//Download zip
			err := downloadZipFile(zipFiles[i], downloadProgress)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
				return
			}

			//Unzip csv files
			zipFiles[i].csvFiles, err = unzipFile(zipFiles[i], dest, unzipProgress)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
				return
			}

			//Remove zip
			err = os.Remove(zipFiles[i].path)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
				return
			}
		}(i)
	}

	wg.Add(1)

	go func() {
		defer wg.Done()
		downloadResults := map[string]float64{}
		unzipResults := map[string]float64{}

		for {
			select {
			case dp := <-downloadProgress:
				for k, v := range dp {
					downloadResults[k] = v
				}
			case up := <-unzipProgress:
				for k, v := range up {
					unzipResults[k] = v
				}
			default:
			}

			//Download progress
			var totalDownloadProgress float64
			for _, v := range downloadResults {
				totalDownloadProgress += v
			}
			totalDownloadProgress = totalDownloadProgress / float64(len(downloadResults))

			//Unzip progress
			var totalUnzipProgress float64
			for _, v := range unzipResults {
				totalUnzipProgress += v
			}
			totalUnzipProgress = totalUnzipProgress / float64(len(downloadResults))

			fmt.Printf("\rDownload progress: %.1f%% - Unzip progress: %.1f%%", totalDownloadProgress, totalUnzipProgress)

			//End
			if totalDownloadProgress >= 100 && totalUnzipProgress >= 100 {
				close(downloadProgress)
				close(unzipProgress)
				fmt.Printf("\nCompleted\n")
				break
			}
		}

	}()

	wg.Wait()
}
