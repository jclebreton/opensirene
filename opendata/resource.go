package opendata

import (
	"io"
	"path/filepath"
)

// Resource describes a single resource in the dataset API
type Resource struct {
	Checksum     Checksum        `json:"checksum"`
	CreatedAt    string          `json:"created_at"`
	Description  *string         `json:"description"`
	Extras       Extras          `json:"extras"`
	Filesize     *int            `json:"filesize"`
	Filetype     string          `json:"filetype"`
	Format       string          `json:"format"`
	ID           string          `json:"id"`
	IsAvailable  bool            `json:"is_available"`
	LastModified string          `json:"last_modified"`
	Latest       string          `json:"latest"`
	Metrics      ResourceMetrics `json:"metrics"`
	Mime         *string         `json:"mime"`
	Published    string          `json:"published"`
	Title        string          `json:"title"`
	URL          string          `json:"url"`
}

func (r Resource) Download(to io.Writer) {

}

func (r Resource) DownloadToFile(to string) {
	filepath.Split(to)
}
