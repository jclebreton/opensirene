package logic

import (
	"reflect"
	"testing"

	"github.com/jclebreton/opensirene/opendata/gouvfr/sirene"
)

func TestRemoteFiles_Diff(t *testing.T) {
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
			if got := Diff(tt.args, tt.rfs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoteFiles.Diff() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoteFiles_ToCSVImport(t *testing.T) {
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
			got, err := ToCSVImport(tt.rfs)
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
