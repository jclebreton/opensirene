package download_extract

import (
	"fmt"
	"io"
)

type Progression struct {
	Name string
	Curr float64
}

type ProgressionBar struct {
	Progress chan Progression
	Total    int
}

//Progress will print the progression percents to stdout
func Progress(downloadProgress, unzipProgress, importProgress, updateProgress ProgressionBar) {
	downloadList := make(map[string]float64)
	unzipList := make(map[string]float64)
	importList := make(map[string]float64)
	updateList := make(map[string]float64)

	for {
		select {
		case p := <-downloadProgress.Progress:
			downloadList[p.Name] = p.Curr
		case p := <-unzipProgress.Progress:
			unzipList[p.Name] = p.Curr
		case p := <-importProgress.Progress:
			importList[p.Name] = p.Curr
		case p := <-updateProgress.Progress:
			updateList[p.Name] = p.Curr
		default:
		}

		totalDownload := getPercentTotal(downloadList, downloadProgress.Total)
		totalUnzip := getPercentTotal(unzipList, unzipProgress.Total)
		totalImport := getPercentTotal(importList, importProgress.Total)
		totalUpdate := getPercentTotal(updateList, updateProgress.Total)

		fmt.Printf("\rDownload: %.2f%% - Unzip: %.2f%% - Import: %.2f%% - Update: %.2f%%",
			totalDownload,
			totalUnzip,
			totalImport,
			totalUpdate,
		)

		if totalDownload >= 100 && totalUnzip >= 100 && totalImport >= 100 && totalUpdate >= 100 {
			break
		}
	}
}

func getPercentTotal(list map[string]float64, total int) float64 {
	if len(list) == 0 {
		return 0
	}

	var t float64
	for _, v := range list {
		t += v
	}
	return t / float64(total)
}

// PassThru code originally from
// http://stackoverflow.com/a/22422650/613575
type PassThru struct {
	io.Reader
	curr     int64
	total    float64
	filename string
	progress chan Progression
}

//Override native Read
func (pt *PassThru) Read(p []byte) (int, error) {
	n, err := pt.Reader.Read(p)
	pt.curr += int64(n)

	// last read will have EOF err
	if err == nil || (err == io.EOF && n > 0) {
		pt.progress <- Progression{Name: pt.filename, Curr: (float64(pt.curr) / pt.total) * 100}
	}

	return n, err
}
