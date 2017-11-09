package domain

import "time"

// UpdateFileStatus is a struct mapping update files saved into the db
type UpdateFileStatus struct {
	ID        int32
	Datetime  time.Time
	Filename  string
	IsSuccess bool
	Err       string
}

// Health is a struct mapping service data for monitoring
type Health struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}
