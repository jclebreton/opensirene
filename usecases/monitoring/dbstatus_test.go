package monitoring

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	d "github.com/jclebreton/opensirene/domain"
	"github.com/jclebreton/opensirene/interfaces/storage/db_status/mocked"
	"github.com/jclebreton/opensirene/testHelpers"
	"github.com/pkg/errors"
)

func TestFindHistories(t *testing.T) {
	tests := []struct {
		retH   []d.UpdateFileStatus
		retErr error
		exp    []d.UpdateFileStatus
		pass   bool
	}{
		{retErr: errors.New("hey"), pass: false},
		{retH: nil, pass: false},
		{retH: []d.UpdateFileStatus{}, pass: false},
		{retH: []d.UpdateFileStatus{d.UpdateFileStatus{ID: 1}}, exp: []d.UpdateFileStatus{d.UpdateFileStatus{ID: 1}}, pass: true},
	}
	for k, tt := range tests {
		r := GetDBStatusRequest{}
		i := &Interactor{MonitoringRW: mocked.RW{FindDatabaseStatusRet: mocked.FindDatabaseStatusRet{tt.retH, tt.retErr}}}

		returned, err := r.findDBSatus(i)

		if tt.pass {
			assert.NoError(t, err, fmt.Sprintf("test %d should pass", k))
			assert.NoError(t, testHelpers.CompareUpdateStatusSlices(tt.exp, returned))
		} else {
			assert.Error(t, err, fmt.Sprintf("test %d should fail", k))
		}

	}
}
