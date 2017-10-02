package download_extract

import (
	"fmt"
	"io"
)

//Progress will print the progression percents to stdout
func Progress(nbZipFiles int, downloadProgress <-chan map[string]float64, unzipProgress <-chan map[string]float64, importProgress <-chan map[string]float64) {
	downloadResults := map[string]float64{}
	unzipResults := map[string]float64{}
	importResults := map[string]float64{}

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
		case up := <-importProgress:
			for k, v := range up {
				importResults[k] = v
			}
		default:
		}

		//Download Progress
		var totalDownloadProgress float64
		for _, v := range downloadResults {
			totalDownloadProgress += v
		}
		totalDownloadProgress = totalDownloadProgress / float64(nbZipFiles)

		//Unzip Progress
		var totalUnzipProgress float64
		for _, v := range unzipResults {
			totalUnzipProgress += v
		}
		totalUnzipProgress = totalUnzipProgress / float64(nbZipFiles)

		//Unzip Progress
		var totalImportProgress float64
		for _, v := range importResults {
			totalImportProgress += v
		}
		//totalImportProgress = totalImportProgress / float64(nbZipFiles)

		fmt.Printf("\rDownload progress: %.2f%% - Unzip progress: %.2f%% - Import progress: %.2f%%", totalDownloadProgress, totalUnzipProgress, totalImportProgress)

		//End
		if totalDownloadProgress >= 100 && totalUnzipProgress >= 100 && totalImportProgress >= 100 {
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
