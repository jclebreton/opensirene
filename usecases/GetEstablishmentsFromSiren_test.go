package usecases

import (
	"errors"
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	d "github.com/jclebreton/opensirene/domain"
	"github.com/jclebreton/opensirene/interfaces/storage/establishments/mocked"
)

func TestGetEstablishmentsFromSiren(t *testing.T) {
	tests := []struct {
		ret  *[]d.Establishment
		exp  *[]d.Establishment
		err  error
		pass bool
	}{
		{err: errors.New("hey"), pass: false},
		{ret: nil, pass: false},
		{ret: &[]d.Establishment{d.Establishment{Siren: "foo"}}, exp: &[]d.Establishment{d.Establishment{Siren: "foo"}}, pass: true},
	}
	for k, tt := range tests {
		r := GetEstablishmentsFromSirenRequest{}
		i := &Interactor{EnterprisesRW: mocked.RW{FindEstablishmentsFromSirenRet: mocked.FindEstablishmentsFromSirenRet{tt.ret, tt.err}}}
		returned, err := r.findEstablishmentsFromSiren(i)

		if tt.pass {
			assert.NoError(t, err, fmt.Sprintf("test %d should pass", k))
			assert.Equal(t, tt.exp, returned)
		} else {
			assert.Error(t, err, fmt.Sprintf("test %d should fail", k))
		}

	}
}
