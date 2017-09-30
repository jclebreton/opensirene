package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_unzip_success(t *testing.T) {
	//Get valid list
	zipFiles, err := getScratchZipList("Jan", url, "/tmp")
	assert.NoError(t, err)

	//Download one file
	progressChan := make(chan map[string]float64, buffer)
	errorsChan := make(chan error, buffer)
	err = downloadZipFile(zipFiles[1], progressChan, errorsChan)
	assert.NoError(t, err)

	//Unzip
	progress := make(chan map[string]float64, buffer)
	csvFiles, err := unzipFile(zipFiles[1], progress)
	assert.NoError(t, err)
	assert.True(t, len(csvFiles) >= 1)
	assert.Regexp(t, "/tmp/sirc-[0-9EQ_]+.csv", csvFiles[0].filename)

	//Check file on disk
	stat, err := os.Stat(csvFiles[0].path)
	assert.NoError(t, err)
	assert.Regexp(t, "sirc-[0-9EQ_]+.csv", stat.Name())
}
