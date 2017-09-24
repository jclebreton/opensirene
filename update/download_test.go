package main

import (
	"testing"

	"os"

	"github.com/stretchr/testify/assert"
)

var buffer = 1000

func Test_downloadZipFile_not_found(t *testing.T) {
	file := zipFile{url: "http://ovh.net/files/notfound.dat"}
	progress := make(chan map[string]float64, buffer)
	errorsChan := make(chan error, buffer)
	err := downloadZipFile(file, progress, errorsChan)
	assert.Error(t, err)
}

func Test_downloadZipFile_success(t *testing.T) {
	file := zipFile{
		filename: "1Mio.dat",
		url:      "http://ovh.net/files/1Mio.dat",
		path:     "/tmp/1Mio.dat",
	}
	progressChan := make(chan map[string]float64, buffer)
	errorsChan := make(chan error, buffer)

	//Download file
	err := downloadZipFile(file, progressChan, errorsChan)
	assert.NoError(t, err)

	//Downloaded file
	stat, err := os.Stat(file.path)
	assert.NoError(t, err)
	assert.Equal(t, int64(1048576), stat.Size())
	assert.Equal(t, "1Mio.dat", stat.Name())

	//Progress chan
	progress := <-progressChan
	assert.True(t, progress["1Mio.dat"] > 0)

	////Errors chan
	//error := <-errorsChan
	//assert.NoError(t, error)
}
