package opendata

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/cheggaaa/pb"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/Depado/lightsiren/conf"
)

var location *time.Location

const (
	nafID           = "59593c53a3a7291dcf9c82bf/"
	sirenID         = "5862206588ee38254d3f4e5e"
	datasetEndpoint = "https://www.data.gouv.fr/api/1/datasets/"
)

// FileType is the main representation for the filetype
type FileType int

// Defines the different filetype constants
const (
	StockType FileType = iota
	DailyType
	MonthlyType
	OtherType
)

// SireneFile is a struct that adds and remove some fields from a Resource
// struct and actually keep only useful fields
type SireneFile struct {
	Checksum Checksum
	URL      string
	FileName string
	Type     FileType
	YearDay  int
	Size     int64
}

// SireneFiles is a slice of pointers to SireneFile
type SireneFiles []*SireneFile

// GetFileSize will make an HEAD request on the file URL to get its size
func (sf *SireneFile) GetFileSize() error {
	var err error
	var resp *http.Response
	var size int
	if resp, err = http.Head(sf.URL); err != nil {
		return errors.Wrapf(err, "can't get %s", sf.URL)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New("Server return non-200 status")
	}
	if size, err = strconv.Atoi(resp.Header.Get("Content-Length")); err != nil {
		return errors.Wrap(err, "can't find size")
	}
	sf.Size = int64(size)
	return nil
}

// DownloadWithProgress will download the file and update the given progress bar
func (sf *SireneFile) DownloadWithProgress(b *pb.ProgressBar) {
	b.Prefix("Downloading " + sf.FileName)
	resp, err := http.Get(sf.URL)
	if err != nil {
		fmt.Printf("Can't get %s: %v\n", sf.URL, err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Server return non-200 status: %v\n", resp.Status)
		return
	}
	dest, err := os.Create(filepath.Join(conf.C.DownloadPath, sf.FileName))
	if err != nil {
		return
	}
	reader := b.NewProxyReader(resp.Body)

	if _, err = io.Copy(dest, reader); err != nil {
		return
	}
}

// UnzipWithProgress will un-compress a zip archive moving all files and folders
// to an output directory
func (sf SireneFile) UnzipWithProgress(b *pb.ProgressBar) error {
	var filenames []string
	b.Prefix("Extracting " + sf.FileName)
	r, err := zip.OpenReader(filepath.Join(conf.C.DownloadPath, sf.FileName))
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		// Store filename/path for returning and using later on
		fpath := filepath.Join(conf.C.DownloadPath, f.Name)
		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
		} else {
			var fdir string
			if lastIndex := strings.LastIndex(fpath, string(os.PathSeparator)); lastIndex > -1 {
				fdir = fpath[:lastIndex]
			}

			err = os.MkdirAll(fdir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
				return err
			}
			f, err := os.OpenFile(
				fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer f.Close()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}

		}
	}
	b.Prefix("Done " + sf.FileName)
	return nil
}

// Grab can be used to grab the full dataset object
func Grab() (*Dataset, error) {
	r, err := http.Get(datasetEndpoint + sirenID)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	target := new(Dataset)

	if err = json.NewDecoder(r.Body).Decode(target); err != nil {
		return target, err
	}
	return target, nil
}

func GrabLatestFull() (SireneFiles, error) {
	var err error
	var ds *Dataset
	var sf *SireneFile
	var sfs SireneFiles

	if ds, err = Grab(); err != nil {
		return sfs, err
	}

	y, m, _ := time.Now().Date()
	first := time.Date(y, m, 1, 0, 0, 0, 0, location).YearDay()

	for _, r := range ds.Resources {
		if sf, err = r.ToSireneFile(); err != nil {
			logrus.WithError(err).Warn("Unprocessable entity")
			continue
		}
		if sf.Type == DailyType && sf.YearDay < first {
			logrus.Warn("Ignored daily that is before the first day of month")
		}
		sfs = append(sfs, sf)
		if sf.Type == StockType {
			break
		}
	}

	return sfs, err
}

func init() {
	var err error
	if location, err = time.LoadLocation("Europe/Paris"); err != nil {
		logrus.WithError(err).Fatal("Couldn't load timezone")
	}
}
