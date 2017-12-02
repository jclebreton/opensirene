package gouvfr

import (
	"net/url"
	ospath "path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/jclebreton/opensirene/domain"
	"github.com/pkg/errors"
)

const SireneID = "5862206588ee38254d3f4e5e"

type SireneR struct{}

func (r *SireneR) GetRemoteFiles(downloadPath string) ([]domain.RemoteFile, error) {
	dataset, err := callGovAPI(SireneID)
	if err != nil {
		return nil, errors.Wrap(err, "unable to call sirene api")
	}

	rf, err := getMonthlyFiles(dataset.Resources, downloadPath)
	if err != nil {
		return nil, errors.Wrap(err, "unable to grab remote files")
	}

	return rf, nil
}

// getRemoteFiles retrieves all the files that needs to be downloaded and
// applied to the database in inverse order (stock first, then each daily file)
func getMonthlyFiles(resources resources, dPath string) ([]domain.RemoteFile, error) {
	var rf *domain.RemoteFile
	var rfs []domain.RemoteFile
	var err error

	y, m, _ := time.Now().Date()
	first := time.Date(y, m, 1, 0, 0, 0, 0, location).YearDay()

	for _, r := range resources {
		if rf, err = newFromResource(r, dPath); err != nil {
			return nil, err
		}

		if rf.Type == domain.DailyType && rf.YearDay < first {
			continue
		}

		rfs = append(rfs, *rf)
		if rf.Type == domain.StockType {
			break
		}
	}

	// Revert the slice to get the right order
	for i, j := 0, len(rfs)-1; i < j; i, j = i+1, j-1 {
		rfs[i], rfs[j] = rfs[j], rfs[i]
	}

	return rfs, err
}

// newFromResource takes an Resource and transforms it to a RemoteFile
func newFromResource(r resource, dPath string) (*domain.RemoteFile, error) {
	var err error
	var u *url.URL
	rf := &domain.RemoteFile{Checksum: domain.Checksum(r.Checksum), URL: r.URL}

	// URL
	if u, err = url.Parse(r.URL); err != nil {
		return nil, err
	}
	base := ospath.Base(u.Path)
	rf.FileName = base
	rf.Path = filepath.Join(dPath, rf.FileName)

	// File type
	bits := strings.SplitN(base, "_", 3)
	if len(bits) < 3 {
		rf.Type = domain.OtherType
		return nil, errors.New("unknown file type")
	}

	switch bits[2] {
	case "E_Q.zip":
		rf.Type = domain.DailyType
	case "L_M.zip":
		rf.Type = domain.StockType
	case "M_M.zip":
		rf.Type = domain.MonthlyType
	}

	// Yearday
	if rf.Type == domain.DailyType {
		yd, err := strconv.Atoi(bits[1][4:])
		if err != nil {
			return nil, errors.Wrap(err, "unparsable day of year")
		}
		rf.YearDay = yd
	}

	return rf, nil
}
