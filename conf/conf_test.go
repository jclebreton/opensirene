package conf

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func Test_Parse_success(t *testing.T) {
	db := Database{
		User:     "foo",
		Name:     "bar",
		Password: "qwerty1234",
		Host:     "127.0.0.1",
		Port:     5432,
		SSLMode:  "disable",
	}
	cors := Cors{
		AllowOrigins:   []string{"foo", "bar"},
		PermissiveMode: true,
	}
	server := Server{
		Host:  "127.0.0.1",
		Port:  8080,
		Debug: true,
		Cors:  cors,
	}
	conf := &Conf{
		Database:     db,
		Server:       server,
		LogLevel:     "fatal",
		DownloadPath: "downloads",
	}

	err := conf.Parse()
	assert.NoError(t, err)
	assert.Equal(t, conf.LogLevel, logrus.GetLevel().String())

	stat, err := os.Stat(conf.DownloadPath)
	assert.NoError(t, err)
	assert.True(t, stat.IsDir())
}
