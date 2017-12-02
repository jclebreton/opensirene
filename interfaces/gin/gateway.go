package gin

import "github.com/jclebreton/opensirene/usecases"

type HttpGateway struct {
	i usecases.Interactor
}

func NewHttpGateway(i usecases.Interactor) HttpGateway {
	return HttpGateway{
		i,
	}
}
