package download_extract

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Passthrue_Read(t *testing.T) {
	//Download temp file
	file := ZipFile{
		filename: "1Mio.dat",
		url:      "http://ovh.net/files/1Mio.dat",
		path:     "/tmp/1Mio.dat",
	}
	progressChan := make(chan map[string]float64, 100000)
	errorsChan := make(chan error, 100000)
	err := file.download(progressChan, errorsChan)
	assert.NoError(t, err)

	//Open it
	r, err := os.Open(file.path)
	assert.NoError(t, err)
	defer r.Close()

	//Read it
	progressChan = make(chan map[string]float64, 100000)
	pt := PassThru{
		Reader:   r,
		filename: "foo",
		progress: progressChan,
	}
	p := []byte{1, 12, 2}
	read, err := pt.Read(p)

	assert.NoError(t, err)
	assert.Equal(t, 3, read)

	//Progress
	progress := <-progressChan
	assert.True(t, progress["foo"] > 0)
}
