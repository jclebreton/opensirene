package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

//downloadZipFile will download a file
func downloadZipFile(file zipFile, progress chan map[string]float64, errorsChan chan error) error {
	resp, _ := http.Get(file.url)
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err := fmt.Errorf("Remote file not found: %s", file.filename)
		progress <- map[string]float64{file.filename: 100}
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
