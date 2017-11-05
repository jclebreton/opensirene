package logic

import (
	"reflect"
	"testing"

	"github.com/jclebreton/opensirene/opendata/gouvfr/sirene"
)

func Test_toCSVImport(t *testing.T) {
	generic := sirene.RemoteFile{
		FileName:       "test.zip",
		ExtractedFiles: []string{"test.csv"},
		Type:           sirene.DailyType,
	}
	genericout := sirene.CSVImport{
		Path:    "test.csv",
		Kind:    sirene.DailyType,
		ZipName: "test.zip",
	}
	tests := []struct {
		name    string
		rfs     sirene.RemoteFiles
		want    CSVImports
		wantErr bool
	}{
		{"should work", sirene.RemoteFiles{&generic}, CSVImports{&genericout}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := toCSVImport(tt.rfs)
			if (err != nil) != tt.wantErr {
				t.Errorf("RemoteFiles.ToCSVImport() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoteFiles.ToCSVImport() = %v, want %v", got, tt.want)
			}
		})
	}
}
