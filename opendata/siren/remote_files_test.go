package siren

import (
	"reflect"
	"testing"
)

func TestRemoteFiles_Diff(t *testing.T) {
	withtest := RemoteFiles{&RemoteFile{FileName: "test.zip"}}
	tests := []struct {
		name string
		rfs  RemoteFiles
		args []string
		want RemoteFiles
	}{
		{"should be empty", withtest, []string{"test.zip"}, RemoteFiles{}},
		{"should not change", withtest, []string{"test2.zip"}, withtest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.rfs.Diff(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoteFiles.Diff() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoteFiles_ToCSVImport(t *testing.T) {
	generic := RemoteFile{
		FileName:       "test.zip",
		ExtractedFiles: []string{"test.csv"},
		Type:           DailyType,
	}
	genericout := CSVImport{
		path:    "test.csv",
		Kind:    DailyType,
		ZipName: "test.zip",
	}
	tests := []struct {
		name    string
		rfs     RemoteFiles
		want    CSVImports
		wantErr bool
	}{
		{"should work", RemoteFiles{&generic}, CSVImports{&genericout}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.rfs.ToCSVImport()
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
