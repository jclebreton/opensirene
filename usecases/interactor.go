package usecases

type Interactor struct {
	DBStatusRW
	EnterprisesRW
	JsonW
	SireneR
}

func NewInteractor(d DBStatusRW, e EnterprisesRW, j JsonW, s SireneR) Interactor {
	return Interactor{
		DBStatusRW:    d,
		EnterprisesRW: e,
		JsonW:         j,
		SireneR:       s,
	}
}
