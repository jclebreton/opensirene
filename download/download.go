package download

import (
	"fmt"
	"time"

	"github.com/jclebreton/opensirene/opendata/siren"
	"github.com/cheggaaa/pb"
)

func worker(id int, bar *pb.ProgressBar, jobs <-chan *siren.RemoteFile, results chan<- error) {
	var err error
	var ok bool

	for s := range jobs {
		bar.Set(0)
		bar.ShowPercent = true
		bar.ShowFinalTime = true
		bar.ShowTimeLeft = true

		// If file is already present on disk, skip the downlod part
		if !s.OnDisk {
			// Download
			bar.Prefix("[Downloading] " + s.FileName)
			if err = s.DownloadWithProgress(bar); err != nil {
				results <- err
				continue
			}

			// Checksum
			bar.Prefix("[Checksum] " + s.FileName)
			if ok, err = s.ChecksumMatch(); err != nil {
				results <- err
				continue
			} else if !ok {
				results <- fmt.Errorf("Checksum did not match for %s", s.FileName)
				continue
			}
		}

		// Extracting
		bar.Prefix("[Unzipping] " + s.FileName)
		if err = s.Unzip(); err != nil {
			results <- err
			continue
		}

		bar.Prefix("[Done] " + s.FileName)
		results <- nil
	}
}

// Do downloads and processes the sirene files
func Do(sfs siren.RemoteFiles, workers int) error {
	var err error
	var pool *pb.Pool

	size := len(sfs)
	jobs := make(chan *siren.RemoteFile, size)
	results := make(chan error, size)

	if pool, err = pb.StartPool(); err != nil {
		return err
	}
	defer pool.Stop()

	for w := 1; w <= workers; w++ {
		bar := pb.New(0).SetUnits(pb.U_BYTES).SetRefreshRate(time.Millisecond * 10)
		pool.Add(bar)
		go worker(w, bar, jobs, results)
	}
	for _, s := range sfs {
		jobs <- s
	}
	close(jobs)

	for i := 1; i <= size; i++ {
		err = <-results
		if err != nil {
			return err
		}
	}
	close(results)
	return nil
}
