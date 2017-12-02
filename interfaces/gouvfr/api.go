package gouvfr

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	host = "https://www.data.gouv.fr"
	path = "/api/1/datasets/"
)

var location *time.Location

func init() {
	var err error
	if location, err = time.LoadLocation("Europe/Paris"); err != nil {
		logrus.WithError(err).Fatal("Couldn't load timezone")
	}
}

// callGovAPI can be used to grab the full dataset object
func callGovAPI(apiIdentifier string) (*dataset, error) {

	url, err := url.Parse(host)
	if err != nil {
		return nil, err
	}
	url.Path = path + apiIdentifier
	r, err := http.Get(url.String())
	if err != nil {
		return nil, err
	}

	if r.StatusCode != http.StatusOK {
		return nil, errors.New(url.String() + " returns a non 200 HTTP status code")
	}

	target := &dataset{}
	defer r.Body.Close()
	if err = json.NewDecoder(r.Body).Decode(target); err != nil {
		return nil, err
	}

	return target, nil
}

// dataset is the top level structure defining a dataset
type dataset struct {
	Badges           []badge        `json:"badges"`
	CreatedAt        string         `json:"created_at"`
	Deleted          *bool          `json:"deleted"`
	Description      string         `json:"description"`
	Extras           *extras        `json:"extras"`
	Frequency        string         `json:"frequency"`
	FrequencyDate    string         `json:"frequency_date"` // Next expected update date, you will be notified once that date is reached.
	ID               string         `json:"id"`
	LastModified     string         `json:"last_modified"`
	LastUpdate       string         `json:"last_update"`
	License          string         `json:"license"`
	Metrics          *metrics       `json:"metrics"`
	Organization     *organization  `json:"organization"`
	Owner            *userReference `json:"owner"`
	Page             string         `json:"page"`
	Private          bool           `json:"private"`
	Resources        resources      `json:"resources"`
	Slug             string         `json:"slug"`
	Spatial          spatial        `json:"spatial"`
	Tags             []string       `json:"tags"`
	TemporalCoverage interface{}    `json:"temporal_coverage"` // TODO: Figure out what this field is
	Title            string         `json:"title"`
	URI              string         `json:"uri"`
}

// resource describes a single resource in the dataset API
type resource struct {
	Checksum     checksum        `json:"checksum"`
	CreatedAt    string          `json:"created_at"`
	Description  *string         `json:"description"`
	Extras       extras          `json:"extras"`
	Filesize     *int            `json:"filesize"`
	Filetype     string          `json:"filetype"`
	Format       string          `json:"format"`
	ID           string          `json:"id"`
	IsAvailable  bool            `json:"is_available"`
	LastModified string          `json:"last_modified"`
	Latest       string          `json:"latest"`
	Metrics      resourceMetrics `json:"metrics"`
	Mime         *string         `json:"mime"`
	Published    string          `json:"published"`
	Title        string          `json:"title"`
	URL          string          `json:"url"`
}

// badge represents a single badge
type badge struct {
	Kind string `json:"kind"`
}

// metrics contains some information about
type metrics struct {
	Badges         int `json:"badges"`
	Discussions    int `json:"discussions"`
	Followers      int `json:"followers"`
	Issues         int `json:"issues"`
	NbHits         int `json:"nb_hits"`
	NbUniqVisitors int `json:"nb_uniq_visitors"`
	NbVisits       int `json:"nb_visits"`
	Reuses         int `json:"reuses"`
	Views          int `json:"views"`
}

// resourceMetrics is a subset of Metrics which only contain Resource metrics
type resourceMetrics struct {
	NbHits         int `json:"nb_hits"`
	NbUniqVisitors int `json:"nb_uniq_visitors"`
	NbVisits       int `json:"nb_visits"`
	Views          int `json:"views"`
}

// organization represents a single org
type organization struct {
	Acronym       interface{} `json:"acronym"`
	Class         string      `json:"class"`
	ID            string      `json:"id"`
	Logo          string      `json:"logo"`
	LogoThumbnail string      `json:"logo_thumbnail"`
	Name          string      `json:"name"`
	Page          string      `json:"page"`
	Slug          string      `json:"slug"`
	URI           string      `json:"uri"`
}

// checksum describes a checksum that can be used to check for file validity
// for example using sha256 in the Type field and the actual sha256 in the Value
type checksum struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// resources is a slice of Resource
type resources []resource

// spatial contains some spatial information
type spatial struct {
	Geom        interface{} `json:"geom"`
	Granularity string      `json:"granularity"`
	Zones       []string    `json:"zones"`
}

// userReference represents a single user
// TODO: Generate that struct
type userReference interface{}

// extras is an unfixed key-value object
type extras interface{}
