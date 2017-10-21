package download

import (
	"sync"
	"time"

	"github.com/Depado/lightsiren/opendata"
	"github.com/cheggaaa/pb"
	"github.com/sirupsen/logrus"
)

func Do(sfs opendata.SireneFiles) {
	var err error
	var wg sync.WaitGroup

	for _, sf := range sfs {
		wg.Add(1)
		go func(s *opendata.SireneFile) {
			defer wg.Done()
			if err = s.GetFileSize(); err != nil {
				logrus.WithError(err).Fatal("Can't get file size")
			}
		}(sf)
	}
	wg.Wait()
	pool, err := pb.StartPool()
	if err != nil {
		panic(err)
	}
	for _, sf := range sfs {
		wg.Add(1)
		go func(s *opendata.SireneFile) {
			defer wg.Done()
			bar := pb.New(int(s.Size)).SetUnits(pb.U_BYTES).SetRefreshRate(time.Millisecond * 10)
			pool.Add(bar)
			s.DownloadWithProgress(bar)
			s.UnzipWithProgress(bar)
		}(sf)
	}
	wg.Wait()
	pool.Stop()
}
