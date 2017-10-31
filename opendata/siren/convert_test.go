package siren

import (
	"reflect"
	"testing"
)

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
