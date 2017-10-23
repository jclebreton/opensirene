package siren

import (
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/Depado/lightsiren/conf"
	"github.com/Depado/lightsiren/opendata"
	"github.com/pkg/errors"
)

// NewFromResource takes an opendata.Resource and transforms it to a
// siren.RemoteFile
func NewFromResource(r opendata.Resource) (*RemoteFile, error) {
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
