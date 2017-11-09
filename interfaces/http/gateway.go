package http

import "github.com/jclebreton/opensirene/usecases"

type HttpGateway struct {
	i usecases.Interactor
}

func NewHttpGateway(hRW usecases.HistoryRW) HttpGateway {
	return HttpGateway{
		i: usecases.Interactor{
			HistoryRW: hRW,
		},
	}
}
