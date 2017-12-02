package monitoring

import "github.com/jclebreton/opensirene/domain"

type GetPingRequest struct{}

// GetPingRequest returns a minimalist string for monitoring
func (i *Interactor) GetPingRequest(r GetPingRequest) domain.Pong {
	return r.getPing(i)
}

func (r *GetPingRequest) getPing(i *Interactor) domain.Pong {
	return "pong"
}
