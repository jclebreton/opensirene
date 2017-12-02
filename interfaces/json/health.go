package json

import "github.com/jclebreton/opensirene/domain"

type formatGetHealthResp struct{}

// HealthJSON is a struct mapping service data for monitoring
type HealthJSON struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func (jw *formatGetHealthResp) FormatHealthResp(h *domain.Health) interface{} {
	return HealthJSON{
		Name:    h.Name,
		Version: h.Version,
	}
}
