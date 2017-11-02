package sirene

import (
	"testing"

	"time"

	"github.com/stretchr/testify/assert"
)

func Test_Conversion_Success(t *testing.T) {
	tests := map[string]string{
		"1954-09-26T10:09:08": "09-26-1954",
		"19540926":            "09-26-1954",
		"195409":              "09-01-1954",
		"1954":                "01-01-1954",
	}
	for in, out := range tests {
		ds, err := NewDateSirene(in)
		assert.NoError(t, err)
		assert.Equal(t, out, ds.String())
		assert.IsType(t, time.Time{}, ds.GetDate())
	}
}

func Test_Conversion_Error(t *testing.T) {
	tests := []string{"1954f-09-26 T 10:09:08", "19540926555"}
	for _, in := range tests {
		ds, err := NewDateSirene(in)
		assert.Error(t, err)
		assert.Equal(t, "01-01-0001", ds.String())
		assert.IsType(t, time.Time{}, ds.GetDate())
	}
}
