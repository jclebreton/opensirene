package usecases

import (
	"github.com/jclebreton/opensirene/domain"
	"github.com/pkg/errors"
)

type GetHistoriesRequest struct {
	// could contain anything useful in order to process the complete usecase
}

func (i Interactor) GetHistories(r GetHistoriesRequest) ([]domain.History, error) {
	// all the business logic happens here ... maybe this case is a bit too simple...
	return r.findHistories(i)
}

// use privates methods in order not to polute the namespace,
// one private method by interactor call so its easy to test
func (r GetHistoriesRequest) findHistories(i Interactor) ([]domain.History, error) {
	hh, err := i.HistoryRW.FindHistories()
	if err != nil {
		return nil, err
	}

	// I don't know if it's really the expected behavior, it's more for testing fun
	if hh == nil || len(hh) == 0 {
		return nil, errors.New("nothing found")
	} else {
		processed := []int32{}
		for _, h := range hh {
			if h.ID == 0 {
				return nil, errors.New("a record with no ID has been returned")
			}
			for _, id := range processed {
				if id == h.ID {
					return nil, errors.New("duplicate records have been returned")
				}
			}

			processed = append(processed, h.ID)
		}
	}

	return hh, nil
}
