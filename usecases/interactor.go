package usecases

type Interactor struct {
	HistoryRW
	JsonW
}

func NewInteractor(hRW HistoryRW, jW JsonW) Interactor {
	return Interactor{
		HistoryRW: hRW,
		JsonW:     jW,
	}
}
