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
