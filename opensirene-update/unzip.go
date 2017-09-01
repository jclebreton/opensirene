package main

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//unzipFile will un-compress a zip archive
func unzipFile(zipFile zipFile, dest string, progress chan map[string]float64) ([]csvFile, error) {

	r, err := zip.OpenReader(zipFile.path)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	totalFiles := len(r.File)
	var result []csvFile

	for k, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return nil, err
		}
		defer rc.Close()

		fpath := filepath.Join(dest, f.Name)

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
				return nil, err
			}
			f, err := os.OpenFile(
				fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return nil, err
			}
			defer f.Close()

			_, err = io.Copy(f, rc)
			if err != nil {
				return nil, err
			}

			result = append(result, csvFile{
				filename: f.Name(),
				path:     fpath,
			})
		}

		progress <- map[string]float64{zipFile.filename: float64(((k + 1) / totalFiles) * 100)}
	}

	return result, nil
}
