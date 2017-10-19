package opendata

const id = "5862206588ee38254d3f4e5e"

type Badge struct {
	Kind string `json:"kind"`
}

type Metrics struct {
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

type Organization struct {
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

type Resource struct {
	Checksum struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"checksum"`
	CreatedAt   string      `json:"created_at"`
	Description interface{} `json:"description"`
	Extras      struct {
	} `json:"extras"`
	Filesize     interface{} `json:"filesize"`
	Filetype     string      `json:"filetype"`
	Format       string      `json:"format"`
	ID           string      `json:"id"`
	IsAvailable  bool        `json:"is_available"`
	LastModified string      `json:"last_modified"`
	Latest       string      `json:"latest"`
	Metrics      struct {
		NbHits         int `json:"nb_hits"`
		NbUniqVisitors int `json:"nb_uniq_visitors"`
		NbVisits       int `json:"nb_visits"`
		Views          int `json:"views"`
	} `json:"metrics"`
	Mime      interface{} `json:"mime"`
	Published string      `json:"published"`
	Title     string      `json:"title"`
	URL       string      `json:"url"`
}

type Resources []Resource

type Spatial struct {
	Geom        interface{} `json:"geom"`
	Granularity string      `json:"granularity"`
	Zones       []string    `json:"zones"`
}

// TODO: Generate that struct
type UserReference interface{}

type Dataset struct {
	Badges           []Badge       `json:"badges"`
	CreatedAt        string        `json:"created_at"`
	Deleted          *bool         `json:"deleted"`
	Description      string        `json:"description"`
	Extras           interface{}   `json:"extras"` // Unfixed key-value pairs
	Frequency        string        `json:"frequency"`
	FrequencyDate    string        `json:"frequency_date"` // Next expected update date, you will be notified once that date is reached.
	ID               string        `json:"id"`
	LastModified     string        `json:"last_modified"`
	LastUpdate       string        `json:"last_update"`
	License          string        `json:"license"`
	Metrics          Metrics       `json:"metrics"`
	Organization     Organization  `json:"organization"`
	Owner            UserReference `json:"owner"`
	Page             string        `json:"page"`
	Private          bool          `json:"private"`
	Resources        Resources     `json:"resources"`
	Slug             string        `json:"slug"`
	Spatial          Spatial       `json:"spatial"`
	Tags             []string      `json:"tags"`
	TemporalCoverage interface{}   `json:"temporal_coverage"`
	Title            string        `json:"title"`
	URI              string        `json:"uri"`
}
