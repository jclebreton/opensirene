package main

import (
	"fmt"
	"log"
	"os"
)

func downloadAndExtract(zipFiles []zipFile, nbWorkers int) ([]csvFile, error) {
	//Progress
	downloadProgressChan := make(chan map[string]float64)
	unzipProgressChan := make(chan map[string]float64)
	go progress(len(zipFiles), downloadProgressChan, unzipProgressChan)
	defer close(downloadProgressChan)
	defer close(unzipProgressChan)

	errorsChan := make(chan error)
	endErrorsChan := make(chan bool)
	defer close(errorsChan)
	defer close(endErrorsChan)
	go listenErrors(errorsChan, endErrorsChan)

	//Workers
	nbZipFiles := len(zipFiles)
	workerChan := make(chan zipFile)
	resultChan := make(chan []csvFile, nbZipFiles)
	defer close(workerChan)
	defer close(resultChan)
	for id := 1; id <= nbWorkers; id++ {
		go startWorker(id, workerChan, resultChan, downloadProgressChan, unzipProgressChan, errorsChan)
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

	endErrorsChan <- true

	return csvFiles, nil
}

func listenErrors(errorsChan chan error, end chan bool) {
	var errors []error
loop:
	for {
		select {
		case error := <-errorsChan:
			errors = append(errors, error)
		case <-end:
			break loop
		default:
		}
	}
	fmt.Printf("\nNumber of errors: %d", len(errors))
	for _, err := range errors {
		fmt.Printf("\n- %s", err)
	}
}

func startWorker(id int, workerChan <-chan zipFile, resultChan chan<- []csvFile, downloadProgressChan,
	unzipProgressChan chan map[string]float64, errorsChan chan error) {
	for zipFile := range workerChan {
		err := downloadZipFile(zipFile, downloadProgressChan, errorsChan)
		if err != nil {
			unzipProgressChan <- map[string]float64{zipFile.filename: 100}
			resultChan <- []csvFile{}
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
