package main

import (
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"time"
)

//getFiles returns a list of files to download
func getZipList(url, dest string) (files []zipFile) {
	currentYear, currentMonth, _ := time.Now().Date()

	//Stock file
	pattern := "sirene_%d%d_L_M"
	if currentMonth-1 < 10 {
		pattern = "sirene_%d0%d_L_M"
	}
	file := zipFile{
		updateType: "complete",
		name:       fmt.Sprintf(pattern, currentYear, currentMonth-1),
	}
	file.filename = file.name + ".zip"
	file.url = fmt.Sprintf("%s/%s", url, file.filename)
	file.path = fmt.Sprintf("%s/%s", dest, file.filename)
	files = append(files, file)

	//Incremental files
	firstDayOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, time.Local)
	firstDayOfYear := time.Date(currentYear, time.January, 1, 0, 0, 0, 0, time.Local)
	startNumber := int(math.Ceil(firstDayOfMonth.Local().Sub(firstDayOfYear).Hours()/24)) + 1
	n := int(math.Ceil(time.Now().Sub(firstDayOfMonth).Hours() / 24))
	for i := 2; i <= n-1; i++ {
		day := time.Date(currentYear, currentMonth, i, 0, 0, 0, 0, time.Local)
		if day.Format("Monday") != "Saturday" && day.Format("Monday") != "Sunday" {
			file = zipFile{
				updateType: "incremental",
				name:       fmt.Sprintf("sirene_%d%d_E_Q", currentYear, startNumber+i-1),
			}
			file.filename = file.name + ".zip"
			file.url = fmt.Sprintf("%s/%s", url, file.filename)
			file.path = fmt.Sprintf("%s/%s", dest, file.filename)
			files = append(files, file)
		}
	}

	return files
}

//downloadZipFile will download a file
func downloadZipFile(file zipFile, progress chan map[string]float64) (err error) {
	resp, _ := http.Get(file.url)
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("Remote file not found: %s", file.filename)
	}

	out, _ := os.Create(file.path)
	defer out.Close()

	src := &PassThru{
		Reader:   resp.Body,
		total:    float64(resp.ContentLength),
		filename: file.filename,
		progress: progress,
	}

	_, err = io.Copy(out, src)
	if err != nil {
		return err
	}

	return nil
}
