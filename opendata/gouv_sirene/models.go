package gouv_sirene

// Badge represents a single badge
type Badge struct {
	Kind string `json:"kind"`
}

// Metrics contains some information about
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

// ResourceMetrics is a subset of Metrics which only contain Resource metrics
type ResourceMetrics struct {
	NbHits         int `json:"nb_hits"`
	NbUniqVisitors int `json:"nb_uniq_visitors"`
	NbVisits       int `json:"nb_visits"`
	Views          int `json:"views"`
}

// Organization represents a single org
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

// Checksum describes a checksum that can be used to check for file validity
// for example using sha256 in the Type field and the actual sha256 in the Value
type Checksum struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// Resources is a slice of Resource
type Resources []Resource

// Spatial contains some spatial information
type Spatial struct {
	Geom        interface{} `json:"geom"`
	Granularity string      `json:"granularity"`
	Zones       []string    `json:"zones"`
}

// UserReference represents a single user
// TODO: Generate that struct
type UserReference interface{}

// Extras is an unfixed key-value object
type Extras interface{}

// Dataset is the top level structure defining a dataset
type Dataset struct {
	Badges           []Badge        `json:"badges"`
	CreatedAt        string         `json:"created_at"`
	Deleted          *bool          `json:"deleted"`
	Description      string         `json:"description"`
	Extras           *Extras        `json:"extras"`
	Frequency        string         `json:"frequency"`
	FrequencyDate    string         `json:"frequency_date"` // Next expected update date, you will be notified once that date is reached.
	ID               string         `json:"id"`
	LastModified     string         `json:"last_modified"`
	LastUpdate       string         `json:"last_update"`
	License          string         `json:"license"`
	Metrics          *Metrics       `json:"metrics"`
	Organization     *Organization  `json:"organization"`
	Owner            *UserReference `json:"owner"`
	Page             string         `json:"page"`
	Private          bool           `json:"private"`
	Resources        Resources      `json:"resources"`
	Slug             string         `json:"slug"`
	Spatial          Spatial        `json:"spatial"`
	Tags             []string       `json:"tags"`
	TemporalCoverage interface{}    `json:"temporal_coverage"` // TODO: Figure out what this field is
	Title            string         `json:"title"`
	URI              string         `json:"uri"`
}
