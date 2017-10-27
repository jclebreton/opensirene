package siren

import (
	"archive/zip"
	"crypto/sha256"
	"fmt"
	"hash"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/cheggaaa/pb"
	"github.com/pkg/errors"

	"github.com/Depado/lightsiren/conf"
	"github.com/Depado/lightsiren/opendata"
)

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

// RemoteFile is a struct that adds and remove some fields from a Resource
// struct and actually keep only useful fields
type RemoteFile struct {
	Checksum       opendata.Checksum
	URL            string
	FileName       string
	Path           string
	Type           FileType
	YearDay        int
	Size           int64
	OnDisk         bool
	ExtractedFiles []string
}

// RemoteFiles is a slice of pointers to RemoteFile
type RemoteFiles []*RemoteFile

// CalculateChecksum generates the checksum of the file using the hasher type
// as defined in the Checksum.Type field
func (rf *RemoteFile) CalculateChecksum() (string, error) {
	var err error
	var f *os.File
	var hasher hash.Hash

	switch rf.Checksum.Type {
	case "sha256":
		hasher = sha256.New()
	default:
		return "", fmt.Errorf("unknown hashing function : %s", rf.Checksum.Type)
	}

	if f, err = os.Open(rf.Path); err != nil {
		return "", err
	}
	defer f.Close()

	if _, err = io.Copy(hasher, f); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hasher.Sum(nil)), nil
}

// ChecksumMatch calculates the checksum of the file on disk and checks if it
// matches the checksum defined in the Checksum.Value field
func (rf *RemoteFile) ChecksumMatch() (bool, error) {
	sum, err := rf.CalculateChecksum()
	if err != nil {
		return false, err
	}
	return sum == rf.Checksum.Value, nil
}

// DownloadWithProgress will download the file and update the given progress bar
func (rf *RemoteFile) DownloadWithProgress(b *pb.ProgressBar) error {
	var err error
	var size int
	var resp *http.Response
	var dest *os.File

	rf.Path = filepath.Join(conf.C.DownloadPath, rf.FileName)

	if resp, err = http.Get(rf.URL); err != nil {
		return errors.Wrapf(err, "couldn't download %s", rf.URL)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("server returned %d for %s", resp.StatusCode, rf.URL)
	}

	if size, err = strconv.Atoi(resp.Header.Get("Content-Length")); err != nil {
		return errors.Wrap(err, "can't find size")
	}
	rf.Size = int64(size)
	b.Total = rf.Size

	if dest, err = os.Create(rf.Path); err != nil {
		return errors.Wrapf(err, "couldn't create file %s", rf.Path)
	}
	reader := b.NewProxyReader(resp.Body)

	if _, err = io.Copy(dest, reader); err != nil {
		return errors.Wrapf(err, "couldn't io.Copy %s", rf.Path)
	}

	return nil
}

// Download silently downloads the file
func (rf *RemoteFile) Download() error {
	var err error
	var resp *http.Response
	var dest *os.File

	rf.Path = filepath.Join(conf.C.DownloadPath, rf.FileName)
	if resp, err = http.Get(rf.URL); err != nil {
		return errors.Wrapf(err, "couldn't download %s", rf.URL)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("server returned %d for %s", resp.StatusCode, rf.URL)
	}

	if dest, err = os.Create(rf.Path); err != nil {
		return errors.Wrapf(err, "couldn't create file %s", rf.Path)
	}

	if _, err = io.Copy(dest, resp.Body); err != nil {
		return errors.Wrapf(err, "couldn't io.Copy %s", rf.Path)
	}

	return nil
}

// Unzip will un-compress a zip archive moving all files and folders
// to an output directory
func (rf *RemoteFile) Unzip() error {
	var filenames []string

	r, err := zip.OpenReader(filepath.Join(conf.C.DownloadPath, rf.FileName))
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
			f, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer f.Close()

			if _, err = io.Copy(f, rc); err != nil {
				return err
			}
		}
	}
	rf.ExtractedFiles = filenames
	return nil
}
