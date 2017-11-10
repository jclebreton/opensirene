package usecases

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	d "github.com/jclebreton/opensirene/domain"
	"github.com/jclebreton/opensirene/interfaces/storage/history/mocked"
	"github.com/jclebreton/opensirene/testHelpers"
	"github.com/pkg/errors"
)

func TestFindHistories(t *testing.T) {
	tests := []struct {
		retH   []d.History
		retErr error
		exp    []d.History
		pass   bool
	}{
		{retErr: errors.New("hey"), pass: false},
		{retH: nil, pass: false},
		{retH: []d.History{}, pass: false},
		{retH: []d.History{d.History{}}, pass: false},
		{retH: []d.History{d.History{ID: 1}, d.History{ID: 1}}, pass: false},

		{retH: []d.History{d.History{ID: 1}}, exp: []d.History{d.History{ID: 1}}, pass: true},
		{retH: []d.History{d.History{ID: 1}, d.History{ID: 2}}, exp: []d.History{d.History{ID: 2}, d.History{ID: 1}}, pass: true},
	}
	for k, tt := range tests {
		r := GetHistoriesRequest{}
		i := Interactor{HistoryRW: mocked.RW{mocked.FindHistoriesRet{Histories: tt.retH, Err: tt.retErr}}}

		returned, err := r.findHistories(i)

		if tt.pass {
			assert.NoError(t, err, fmt.Sprintf("test %d should pass", k))
			assert.NoError(t, testHelpers.CompareHistorySlices(tt.exp, returned))
		} else {
			assert.Error(t, err, fmt.Sprintf("test %d should fail", k))
		}

	}
}
