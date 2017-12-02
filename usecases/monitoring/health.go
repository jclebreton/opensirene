package monitoring

import "github.com/jclebreton/opensirene/domain"

type GetHealthRequest struct{}

// GetHealth returns the service status
func (i *Interactor) GetHealth(r GetHealthRequest) *domain.Health {
	return r.getHealth(i)
}

func (r *GetHealthRequest) getHealth(i *Interactor) *domain.Health {
	return i.MonitoringRW.GetHealth()
}
