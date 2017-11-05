package logic

import (
	"reflect"
	"testing"

	"github.com/jclebreton/opensirene/database"
	"github.com/jclebreton/opensirene/opendata/gouvfr/sirene"
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
