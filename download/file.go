package download

import "time"

// Type Constants
const (
	Stock = iota
	Daily
	Monthly
)

// File is a simple struct representing how files are displayed on the download
// page of the siren site
type File struct {
	Name string
	Link string
	Date time.Time
}

// Files is a slice of File used to sort by date
type Files []*File

func (f Files) Len() int {
	return len(f)
}

func (f Files) Less(i, j int) bool {
	return f[i].Date.Before(f[j].Date)
}

func (f Files) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}
