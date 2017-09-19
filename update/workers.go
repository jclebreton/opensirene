package main

import (
	"log"
	"os"
)

func downloadAndExtract(dest string) []csvFile {
	zipFiles := getZipList(url, dest)

	//Progress
	downloadProgressChan := make(chan map[string]float64)
	unzipProgressChan := make(chan map[string]float64)
	go progress(len(zipFiles), downloadProgressChan, unzipProgressChan)
	defer close(downloadProgressChan)
	defer close(unzipProgressChan)

	//Workers
	nbZipFiles := len(zipFiles)
	workerChan := make(chan zipFile)
	resultChan := make(chan []csvFile, nbZipFiles)
	defer close(workerChan)
	defer close(resultChan)
	for id := 1; id <= nbWorkers; id++ {
		go startWorker(id, workerChan, resultChan, downloadProgressChan, unzipProgressChan)
	}

	//Send Zip files
	go func() {
		for _, zipFile := range zipFiles {
			workerChan <- zipFile
		}
	}()

	//Waiting CSV files
	var csvFiles []csvFile
	for i := 1; i <= nbZipFiles; i++ {
		files := <-resultChan
		for _, f := range files {
			csvFiles = append(csvFiles, f)
		}
	}

	return csvFiles
}

func startWorker(id int, workerChan <-chan zipFile, resultChan chan<- []csvFile, downloadProgressChan, unzipProgressChan chan map[string]float64) {
	for zipFile := range workerChan {
		err := downloadZipFile(zipFile, downloadProgressChan)
		if err != nil {
			log.Fatal(err)
			return
		}

		zipFile.csvFiles, err = unzipFile(zipFile, unzipProgressChan)
		if err != nil {
			log.Fatal(err)
			return
		}

		err = os.Remove(zipFile.path)
		if err != nil {
			log.Fatal(err)
			return
		}

		resultChan <- zipFile.csvFiles
	}
}
