package logic

import (
	"os"
	"reflect"
	"testing"

	"github.com/jclebreton/opensirene/conf"
	"github.com/jclebreton/opensirene/database"
	"github.com/jclebreton/opensirene/opendata/gouvfr/sirene"
	"github.com/stretchr/testify/assert"
)

func Test_getFilesToImport(t *testing.T) {
	withtest := sirene.RemoteFiles{&sirene.RemoteFile{FileName: "test.zip"}}
	tests := []struct {
		name string
		rfs  sirene.RemoteFiles
		args []string
		want sirene.RemoteFiles
	}{
		{"should be empty", withtest, []string{"test.zip"}, sirene.RemoteFiles{}},
		{"should not change", withtest, []string{"test2.zip"}, withtest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			crontab := &Crontab{PgxClient: &database.PgxClient{}}
			if got := crontab.getFilesToImport(tt.args, tt.rfs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoteFiles.Diff() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeUselessFiles(t *testing.T) {
	var err error
	var path = "/tmp/test_opensirene/"

	err = os.MkdirAll(path, 0777)
	defer os.RemoveAll(path)
	assert.NoError(t, err)
	file1, err := os.Create(path + "foo")
	assert.NoError(t, err)
	file2, err := os.Create(path + "bar")
	assert.NoError(t, err)

	crontab := &Crontab{Config: conf.Crontab{DownloadPath: path}}
	err = crontab.removeUselessFiles([]string{"foo"})

	_, err = os.Stat(file1.Name())
	assert.True(t, !os.IsNotExist(err))

	_, err = os.Stat(file2.Name())
	assert.True(t, os.IsNotExist(err))
}
