package opendata

import (
	"io"
	"net/url"
	"path"
	"path/filepath"
	"strconv"
	"strings"

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

func (r Resource) Download(to io.Writer) {

}

func (r Resource) DownloadToFile(to string) {
	filepath.Split(to)
}

func (r Resource) ToSireneFile() (*SireneFile, error) {
	var err error
	var u *url.URL

	sf := SireneFile{
		Checksum: r.Checksum,
		URL:      r.URL,
	}

	if u, err = url.Parse(r.URL); err != nil {
		return &sf, err
	}
	base := path.Base(u.Path)
	sf.FileName = base
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
	return &sf, nil
}
