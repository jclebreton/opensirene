package siren

import (
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/jclebreton/opensirene/conf"
	"github.com/pkg/errors"
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

// NewFromResource takes an Resource and transforms it to a
// siren.RemoteFile
func NewFromResource(r Resource) (*RemoteFile, error) {
	var err error
	var u *url.URL

	sf := RemoteFile{
		Checksum: r.Checksum,
		URL:      r.URL,
	}

	if u, err = url.Parse(r.URL); err != nil {
		return &sf, err
	}

	base := path.Base(u.Path)
	sf.FileName = base
	sf.Path = filepath.Join(conf.C.DownloadPath, sf.FileName)
	bits := strings.SplitN(base, "_", 3)
	if len(bits) < 3 {
		sf.Type = OtherType
		return &sf, nil
	}
	switch bits[2] {
	case "E_Q.zip":
		sf.Type = DailyType
	case "L_M.zip":
		sf.Type = StockType
	case "M_M.zip":
		sf.Type = MonthlyType
	}
	if sf.Type == DailyType {
		yd, err := strconv.Atoi(bits[1][4:])
		if err != nil {
			return &sf, errors.Wrap(err, "unparsable day of year")
		}
		sf.YearDay = yd
	}
	if _, err := os.Stat(sf.Path); os.IsNotExist(err) {
		sf.OnDisk = false
	} else {
		sf.OnDisk, err = sf.ChecksumMatch()
		if err != nil {
			return &sf, err
		}
	}
	return &sf, nil
}
