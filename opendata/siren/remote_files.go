package siren

// RemoteFiles is a slice of pointers to RemoteFile
type RemoteFiles []*RemoteFile

// Diff is a function used to filter a slice of RemoteFile with a slice of
// string. If the file name is present in the RemoteFiles, they are evicted from
// the returned slice
func (rfs RemoteFiles) Diff(in []string) RemoteFiles {
	out := RemoteFiles{}
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
func (rfs RemoteFiles) ToCSVImport() (CSVImports, error) {
	var out CSVImports
	for _, rf := range rfs {
		out = append(out, &CSVImport{
			path:    rf.ExtractedFiles[0],
			Kind:    rf.Type,
			ZipName: rf.FileName,
		})
	}
	return out, nil
}
