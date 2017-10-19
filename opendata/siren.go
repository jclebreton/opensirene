package opendata

import (
	"encoding/json"
	"net/http"
)

const (
	sirenID         = "5862206588ee38254d3f4e5e"
	datasetEndpoint = "https://www.data.gouv.fr/api/1/datasets/"
)

// Grab is an experiment
func Grab() error {
	r, err := http.Get(datasetEndpoint + sirenID)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	target := new(Dataset)

	if err = json.NewDecoder(r.Body).Decode(target); err != nil {
		return err
	}
	return nil
}
