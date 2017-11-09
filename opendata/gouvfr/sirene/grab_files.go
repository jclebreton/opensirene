package sirene

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func worker(id int, jobs <-chan *RemoteFile, results chan<- error, dPath string) {
	var err error
	var ok bool

	for s := range jobs {

		// If file is already present on disk, skip the download part
		if !s.OnDisk {
			// Download
			logrus.WithField("file", s.FileName).Info("Start grabbing")
			if err = s.Download(dPath); err != nil {
				results <- err
				continue
			}
		}

		// Checksum
		logrus.WithField("file", s.FileName).Info("Start checksum")
		if ok, err = s.ChecksumMatch(); err != nil {
			results <- err
			continue
		} else if !ok {
			results <- fmt.Errorf("checksum did not match for %s", s.FileName)
			continue
		}

		// Extracting
		logrus.WithField("file", s.FileName).Info("Start unzipping")
		if err = s.Unzip(dPath); err != nil {
			results <- err
			continue
		}

		logrus.WithField("file", s.FileName).Info("Grabbing done")
		results <- nil
	}
}

// Do downloads and processes the sirene files
func Do(sfs RemoteFiles, workers int, dPath string) error {
	var err error

	size := len(sfs)
	jobs := make(chan *RemoteFile, size)
	results := make(chan error, size)

	for w := 1; w <= workers; w++ {
		go worker(w, jobs, results, dPath)
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
