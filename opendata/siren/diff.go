package siren

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
