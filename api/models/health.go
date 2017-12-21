package models

// History is a struct mapping history sql table
type Health struct {
	Name         string                `json:"name"`
	Version      string                `json:"version"`
	Dependencies map[string]Dependency `json:"dependencies"`
}

// Dependency is a struct to describe a dependency
type Dependency struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Error  string `json:"error"`
}
