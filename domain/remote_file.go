package domain

import "os"

// FileType is the main representation for the filetype
type FileType int

// Defines the different filetype constants
const (
	OtherType FileType = iota
	StockType
	DailyType
	MonthlyType
)

// RemoteFile is a struct that adds and remove some fields from a Resource
// struct and actually keep only useful fields
type RemoteFile struct {
	Checksum       Checksum
	URL            string
	FileName       string
	Path           string
	Type           FileType
	YearDay        int
	ExtractedFiles []string
}

// Checksum describes a checksum that can be used to check for file validity
type Checksum struct {
	Type  string
	Value string
}

func (rf *RemoteFile) GetType() string {
	switch rf.Type {
	case StockType:
		return "stock"
	case DailyType:
		return "daily"
	case MonthlyType:
		return "monthly"
	default:
		return "other"
	}
}

func (rf *RemoteFile) IsOnDisk() bool {
	if _, err := os.Stat(rf.Path); err != nil {
		return false
	}
	return true
}
