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
func unzipFile(zipFile zipFile, progress chan map[string]float64) ([]csvFile, error) {

	dest := filepath.Dir(zipFile.path)

	r, err := zip.OpenReader(zipFile.path)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	var result []csvFile

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return nil, err
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
					return nil, err
				}
			}

			f, err := os.OpenFile(
				fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return nil, err
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
				return nil, err
			}

			result = append(result, csvFile{
				filename: file.Name,
				path:     fpath,
			})
		}
	}

	return result, nil
}
