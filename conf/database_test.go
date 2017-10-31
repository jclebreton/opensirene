package conf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ConnectionString(t *testing.T) {
	db := &Database{
		User:     "foo",
		Name:     "bar",
		Password: "qwerty1234",
		Host:     "127.0.0.1",
		Port:     5432,
		SSLMode:  "disable",
	}

	assert.Equal(t, "user=foo password=qwerty1234 dbname=bar host=127.0.0.1 port=5432 sslmode=disable", db.ConnectionString())
}
