package gouv_sirene

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

var location *time.Location

// Grab can be used to grab the full dataset object
func Grab() (*Dataset, error) {
	r, err := http.Get(datasetEndpoint + sirenID)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	target := new(Dataset)

	if err = json.NewDecoder(r.Body).Decode(target); err != nil {
		return target, err
	}
	return target, nil
}

// GrabLatestFull retrieves all the files that needs to be downloaded and
// applied to the database in inverse order (stock first, then each daily file)
func GrabLatestFull() (RemoteFiles, error) {
	var err error
	var ds *Dataset
	var rf *RemoteFile
	var rfs RemoteFiles

	if ds, err = Grab(); err != nil {
		return rfs, err
	}

	y, m, _ := time.Now().Date()
	first := time.Date(y, m, 1, 0, 0, 0, 0, location).YearDay()

	for _, r := range ds.Resources {
		if rf, err = NewFromResource(r); err != nil {
			logrus.WithError(err).Warn("Unprocessable entity")
			continue
		}
		if rf.Type == DailyType && rf.YearDay < first {
			logrus.Warn("Ignored daily that is before the first day of month")
			continue
		}
		rfs = append(rfs, rf)
		if rf.Type == StockType {
			break
		}
	}

	// Revert the slice to get the right order
	for i, j := 0, len(rfs)-1; i < j; i, j = i+1, j-1 {
		rfs[i], rfs[j] = rfs[j], rfs[i]
	}

	return rfs, err
}

func init() {
	var err error
	if location, err = time.LoadLocation("Europe/Paris"); err != nil {
		logrus.WithError(err).Fatal("Couldn't load timezone")
	}
}
