package logic

import (
	"github.com/jclebreton/opensirene/opendata/gouvfr/sirene"
)

// Diff is a function used to filter a slice of RemoteFile with a slice of
// string. If the file name is present in the RemoteFiles, they are evicted from
// the returned slice
func Diff(in []string, rfs sirene.RemoteFiles) sirene.RemoteFiles {
	out := sirene.RemoteFiles{}
	for _, rf := range rfs {
		for _, fn := range in {
			if rf.FileName != fn {
				out = append(out, rf)
			}
		}
	}
	return out
}

// ToCSVImport converts a slice of RemoteFile to a slice of CSVImport.
// It expects that at least one file was extracted
func ToCSVImport(rfs sirene.RemoteFiles) (CSVImports, error) {
	var out CSVImports
	for _, rf := range rfs {
		out = append(out, &sirene.CSVImport{
			Path:    rf.ExtractedFiles[0],
			Kind:    rf.Type,
			ZipName: rf.FileName,
		})
	}
	return out, nil
}
