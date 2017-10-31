package siren

// ToCSVImport converts a slice of RemoteFile to a slice of CSVImport.
// It expects that at least one file was extracted
func (rfs RemoteFiles) ToCSVImport() ([]*CSVImport, error) {
	var out []*CSVImport
	for _, rf := range rfs {
		out = append(out, &CSVImport{
			path:    rf.ExtractedFiles[0],
			Kind:    rf.Type,
			ZipName: rf.FileName,
		})
	}
	return out, nil
}
