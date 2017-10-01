package download_extract

import (
	"fmt"
	"io"
)

//progress will print the progression percents to stdout
func progress(nbZipFiles int, downloadProgress <-chan map[string]float64, unzipProgress <-chan map[string]float64) {
	downloadResults := map[string]float64{}
	unzipResults := map[string]float64{}

	for {
		select {
		case dp := <-downloadProgress:
			for k, v := range dp {
				downloadResults[k] = v
			}
		case up := <-unzipProgress:
			for k, v := range up {
				unzipResults[k] = v
			}
		default:
		}

		//Download progress
		var totalDownloadProgress float64
		for _, v := range downloadResults {
			totalDownloadProgress += v
		}
		totalDownloadProgress = totalDownloadProgress / float64(nbZipFiles)

		//Unzip progress
		var totalUnzipProgress float64
		for _, v := range unzipResults {
			totalUnzipProgress += v
		}
		totalUnzipProgress = totalUnzipProgress / float64(nbZipFiles)

		fmt.Printf("\rDownload progress: %.1f%% - Unzip progress: %.1f%%", totalDownloadProgress, totalUnzipProgress)

		//End
		if totalDownloadProgress >= 100 && totalUnzipProgress >= 100 {
			break
		}
	}
}

// PassThru code originally from
// http://stackoverflow.com/a/22422650/613575
type PassThru struct {
	io.Reader
	curr     int64
	total    float64
	filename string
	progress chan map[string]float64
}

//Override native Read
func (pt *PassThru) Read(p []byte) (int, error) {
	n, err := pt.Reader.Read(p)
	pt.curr += int64(n)

	// last read will have EOF err
	if err == nil || (err == io.EOF && n > 0) {
		pt.progress <- map[string]float64{pt.filename: (float64(pt.curr) / pt.total) * 100}
	}

	return n, err
}
