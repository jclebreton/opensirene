package logic

import (
	"reflect"
	"testing"

	"github.com/jclebreton/opensirene/opendata/gouv_sirene"
)

func TestRemoteFiles_Diff(t *testing.T) {
	withtest := gouv_sirene.RemoteFiles{&gouv_sirene.RemoteFile{FileName: "test.zip"}}
	tests := []struct {
		name string
		rfs  gouv_sirene.RemoteFiles
		args []string
		want gouv_sirene.RemoteFiles
	}{
		{"should be empty", withtest, []string{"test.zip"}, gouv_sirene.RemoteFiles{}},
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
	generic := gouv_sirene.RemoteFile{
		FileName:       "test.zip",
		ExtractedFiles: []string{"test.csv"},
		Type:           gouv_sirene.DailyType,
	}
	genericout := gouv_sirene.CSVImport{
		Path:    "test.csv",
		Kind:    gouv_sirene.DailyType,
		ZipName: "test.zip",
	}
	tests := []struct {
		name    string
		rfs     gouv_sirene.RemoteFiles
		want    CSVImports
		wantErr bool
	}{
		{"should work", gouv_sirene.RemoteFiles{&generic}, CSVImports{&genericout}, false},
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
