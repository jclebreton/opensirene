package usecases

type Interactor struct {
	DBStatusRW
	EnterprisesRW
	JsonW
}

func NewInteractor(hRW DBStatusRW, eRW EnterprisesRW, jW JsonW) Interactor {
	return Interactor{
		DBStatusRW:    hRW,
		EnterprisesRW: eRW,
		JsonW:         jW,
	}
}
