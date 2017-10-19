package opendata

import (
	"encoding/json"
	"net/http"
)

const (
	nafID           = "59593c53a3a7291dcf9c82bf/"
	sirenID         = "5862206588ee38254d3f4e5e"
	datasetEndpoint = "https://www.data.gouv.fr/api/1/datasets/"
)

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
