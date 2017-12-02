package gin

import (
	"github.com/jclebreton/opensirene/usecases"
	"github.com/jclebreton/opensirene/usecases/monitoring"
)

type HttpGateway struct {
	public usecases.Interactor
	admin  monitoring.Interactor
}

func NewHttpGateway(public usecases.Interactor, monitoring monitoring.Interactor) HttpGateway {
	return HttpGateway{
		public,
		monitoring,
	}
}
