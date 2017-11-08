package conf

import (
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/alecthomas/assert"
)

func TestTypes(t *testing.T) {
	// Test that everything goes fine
	type Mystruct struct {
		Int      int           `default:"10"`
		String   string        `default:"str"`
		Bool     bool          `default:"true"`
		UInt     uint          `default:"10"`
		Int64    int64         `default:"10"`
		Float32  float32       `default:"1.10"`
		Float64  float64       `default:"1.10"`
		Duration time.Duration `default:"1h"`
	}
	in := Mystruct{}
	err := Parse(&in)
	assert.NoError(t, err, "error should be nil")
	assert.Equal(t, in.Int, 10)
	assert.Equal(t, in.Int64, int64(10))
	assert.Equal(t, in.String, "str")
	assert.Equal(t, in.Bool, true)
	assert.Equal(t, in.Float32, float32(1.10))
	assert.Equal(t, in.Float64, 1.10)
	assert.Equal(t, in.Duration, time.Hour)

	tests := []struct {
		name  string
		good  interface{}
		wrong interface{}
	}{
		{
			"integer test",
			&struct {
				Int int `default:"10"`
			}{},
			&struct {
				Int int `default:"xxxx"`
			}{},
		},
		{
			"bool test",
			&struct {
				Out bool `default:"true"`
			}{},
			&struct {
				Out bool `default:"xxxx"`
			}{},
		},
		{
			"uint test",
			&struct {
				Out uint `default:"1"`
			}{},
			&struct {
				Out uint `default:"xxxx"`
			}{},
		},
		{
			"int64 test",
			&struct {
				Out int64 `default:"1"`
			}{},
			&struct {
				Out int64 `default:"xxxx"`
			}{},
		},
		{
			"float32 test",
			&struct {
				Out float32 `default:"1.10"`
			}{},
			&struct {
				Out float32 `default:"xxxx"`
			}{},
		},
		{
			"float64 test",
			&struct {
				Out float64 `default:"1.10"`
			}{},
			&struct {
				Out float64 `default:"xxxx"`
			}{},
		},
		{
			"duration test",
			&struct {
				Out time.Duration `default:"1h"`
			}{},
			&struct {
				Out time.Duration `default:"xxxx"`
			}{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Error(t, Parse(tt.wrong))
			assert.NoError(t, Parse(tt.good))
		})
	}
}

func TestParse(t *testing.T) {
	var err error
	type Mystruct struct {
		Myfield int `env:"MYFIELD" default:"10"`
	}
	in := Mystruct{}
	// Test the fallback to the default field tag
	err = Parse(&in)
	assert.NoError(t, err, "error should be nil")
	assert.Equal(t, in.Myfield, 10)

	// Test the environment variable has an impact on the struct
	os.Setenv("MYFIELD", "11")
	err = Parse(&in)
	assert.NoError(t, err, "error should be nil")
	assert.Equal(t, in.Myfield, 11)

	// Test that an error is returned when the type is not compatible
	os.Setenv("MYFIELD", "random")
	err = Parse(&in)
	assert.Error(t, err)

	type erroneous struct {
		Myfield int `default:"x"`
	}
	// Test the fallback to the default field tag
	err = Parse(&erroneous{})
	assert.Error(t, err, "should throw back an error")
}

func Test_isZero(t *testing.T) {
	var nilslice []string
	var nilmap map[string]string
	type args struct {
		v reflect.Value
	}
	tests := []struct {
		name string
		args reflect.Value
		want bool
	}{
		{"nil slice should return true", reflect.ValueOf(nilslice), true},
		{"nil map should return true", reflect.ValueOf(nilmap), true},
		{"slice should be false", reflect.ValueOf([]string{}), false},
		{"map should be false", reflect.ValueOf(new(map[bool]bool)), false},
		{"struct instance should be false", reflect.ValueOf(struct{ One bool }{true}), false},
		{"empty struct should be true", reflect.ValueOf(struct{ One bool }{}), true},
		{"empty array should be true", reflect.ValueOf([0]string{}), true},
		{"array should be false", reflect.ValueOf([1]string{"hi"}), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isZero(tt.args); got != tt.want {
				t.Errorf("isZero() = %v, want %v", got, tt.want)
			}
		})
	}
}
