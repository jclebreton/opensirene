package download_extract

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type ZipFile struct {
	name       string
	filename   string
	path       string
	url        string
	updateType string
	csvFiles   []CsvFile
}

//download will download the zip file
func (file *ZipFile) download(progress chan Progression, errorsChan chan error) error {
	resp, _ := http.Get(file.url)
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err := fmt.Errorf("Remote file not found: %s", file.filename)
		progress <- Progression{Name: file.filename, Curr: 100}
		errorsChan <- err
		return err
	}

	out, _ := os.Create(file.path)
	defer out.Close()

	src := &PassThru{
		Reader:   resp.Body,
		total:    float64(resp.ContentLength),
		filename: file.filename,
		progress: progress,
	}

	_, err := io.Copy(out, src)
	if err != nil {
		errorsChan <- err
		return err
	}

	return nil
}

//remoteFileExist will check if the corresponding remote file exist
func (file *ZipFile) remoteExist() bool {
	resp, _ := http.Head(file.url)
	defer resp.Body.Close()
	return resp.StatusCode == 200
}

//unzip will un-compress the zip archive
func (zipFile *ZipFile) unzip(progress chan Progression) error {

	dest := filepath.Dir(zipFile.path)

	r, err := zip.OpenReader(zipFile.path)
	if err != nil {
		return err
	}
	defer r.Close()

	var result []CsvFile

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		fpath := filepath.Join(dest, f.Name)

		file := f
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
		} else {
			var fdir string
			if lastIndex := strings.LastIndex(fpath, string(os.PathSeparator)); lastIndex > -1 {
				fdir = fpath[:lastIndex]
			}

			if fdir != "" {
				err = os.MkdirAll(fdir, os.ModePerm)
				if err != nil {
					log.Fatal(err)
					return err
				}
			}

			f, err := os.OpenFile(
				fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer f.Close()

			src := &PassThru{
				Reader:   rc,
				total:    float64(file.FileInfo().Size()),
				filename: zipFile.filename,
				progress: progress,
			}

			_, err = io.Copy(f, src)
			if err != nil {
				return err
			}

			result = append(result, CsvFile{
				UpdateType: zipFile.updateType,
				Filename:   file.Name,
				Path:       fpath,
			})
		}
	}

	zipFile.csvFiles = result
	return nil
}
