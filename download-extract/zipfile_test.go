package download_extract

import (
	"testing"

	"os"

	"github.com/stretchr/testify/assert"
)

func Test_downloadZipFile_not_found(t *testing.T) {
	file := ZipFile{url: "http://ovh.net/files/notfound.dat"}
	progress := make(chan map[string]float64, 100000)
	errorsChan := make(chan error, 100000)
	err := file.download(progress, errorsChan)
	assert.Error(t, err)
}

func Test_downloadZipFile_success(t *testing.T) {
	file := ZipFile{
		filename: "1Mio.dat",
		url:      "http://ovh.net/files/1Mio.dat",
		path:     "/tmp/1Mio.dat",
	}
	progressChan := make(chan map[string]float64, 100000)
	errorsChan := make(chan error, 100000)

	//Download file
	err := file.download(progressChan, errorsChan)
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

func Test_unzip_success(t *testing.T) {
	//Get valid list
	zipFiles, err := GetScratchZipList("Jan", "http://files.data.gouv.fr/sirene", "/tmp")
	assert.NoError(t, err)

	//Download one file
	progressChan := make(chan map[string]float64, 100000)
	errorsChan := make(chan error, 100000)
	err = zipFiles[1].download(progressChan, errorsChan)
	assert.NoError(t, err)

	//Unzip
	progress := make(chan map[string]float64, 100000)
	csvFiles, err := zipFiles[1].unzip(progress)
	assert.NoError(t, err)
	assert.True(t, len(csvFiles) >= 1)
	assert.Regexp(t, "sirc-[0-9EQ_]+.csv", csvFiles[0].filename)
	assert.Regexp(t, "/tmp/sirc-[0-9EQ_]+.csv", csvFiles[0].path)

	//Check file on disk
	stat, err := os.Stat(csvFiles[0].path)
	assert.NoError(t, err)
	assert.Regexp(t, "sirc-[0-9EQ_]+.csv", stat.Name())
}
