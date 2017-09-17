package main

import "fmt"

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
