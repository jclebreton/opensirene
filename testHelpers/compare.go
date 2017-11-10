package testHelpers

import (
	"errors"
	"fmt"

	"github.com/jclebreton/opensirene/domain"
)

func CompareHistorySlices(sl1, sl2 []domain.History) error {
	if len(sl1) != len(sl2) {
		return errors.New("slices have != lengths")
	}
	for _, v1 := range sl1 {
		found := false
		for _, v2 := range sl2 {
			if v2.ID == v1.ID {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("Offer %v from slice 1 hasn't been found in slice 2", v1)
		}
	}
	return nil
}
