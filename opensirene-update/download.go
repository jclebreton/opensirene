package main

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

//getIncrementalFiles returns a list of Incremental files to download
func getIncrementalFiles() []string {
	year, month, _ := time.Now().Date()
	firstDayOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	firstDayOfYear := time.Date(year, time.January, 1, 0, 0, 0, 0, time.Local)
	startNumber := int(firstDayOfMonth.Local().Sub(firstDayOfYear).Hours()/24) + 1
	n := int(time.Now().Sub(firstDayOfMonth).Hours() / 24)

	var files []string
	for i := 2; i <= n; i++ {
		day := time.Date(year, month, i, 0, 0, 0, 0, time.Local).Format("Monday")
		if day != "Saturday" && day != "Sunday" {
			files = append(files, fmt.Sprintf("sirene_%d%d_E_Q", year, startNumber+i))
		}
	}

	return files
}

//getLastStockFile returns the name of the last stock file to download
func getLastStockFile() string {
	year, month, _ := time.Now().Date()
	month -= 1
	pattern := "sirene_%d%d_L_M"
	if month < 10 {
		pattern = "sirene_%d0%d_L_M"
	}

	return fmt.Sprintf(pattern, year, month)
}

//downloadFile will download a file
func downloadFile(file, prefix, dest string) (err error) {
	url := fmt.Sprintf("%s/%s", prefix, file)

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	// Stop if file does not exist
	if resp.StatusCode != 200 {
		return errors.New(fmt.Sprintf("%s not found!\n", file))
	}

	// Create the file
	out, err := os.Create(fmt.Sprintf("%s/%s", dest, file))
	defer out.Close()
	if err != nil {
		return err
	}

	// Writer the body to file
	defer resp.Body.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

//unzipFile will un-compress a zip archive
func unzipFile(src, dest string) ([]string, error) {

	var filenames []string

	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer r.Close()

	for _, f := range r.File {

		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}
		defer rc.Close()

		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, f.Name)
		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {

			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)

		} else {

			// Make File
			var fdir string
			if lastIndex := strings.LastIndex(fpath, string(os.PathSeparator)); lastIndex > -1 {
				fdir = fpath[:lastIndex]
			}

			err = os.MkdirAll(fdir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
				return filenames, err
			}
			f, err := os.OpenFile(
				fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return filenames, err
			}
			defer f.Close()

			_, err = io.Copy(f, rc)
			if err != nil {
				return filenames, err
			}

		}
	}
	return filenames, nil
}
