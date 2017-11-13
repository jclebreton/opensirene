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
		Host:   "127.0.0.1",
		Port:   8080,
		Debug:  true,
		Cors:   cors,
		Prefix: Prefix{Api: "/v1", Admin: "/admin"},
	}
	crontab := Crontab{
		DownloadPath: "downloads",
		EveryXHours:  3,
	}
	logger := Logger{
		Level:  "info",
		Format: "text",
	}
	conf := &Conf{
		Database: db,
		Server:   server,
		Crontab:  crontab,
		Logger:   logger,
	}

	err := conf.Parse()
	assert.NoError(t, err)
	assert.Equal(t, logger.Level, logrus.GetLevel().String())

	stat, err := os.Stat(conf.Crontab.DownloadPath)
	assert.NoError(t, err)
	assert.True(t, stat.IsDir())
}

func TestSetLogLevel(t *testing.T) {
	type args struct {
		lvl string
	}
	tests := []struct {
		name     string
		lvl      string
		expected logrus.Level
	}{
		{"must set to debug", "debug", logrus.DebugLevel},
		{"must set to info", "info", logrus.InfoLevel},
		{"must set to warning", "warn", logrus.WarnLevel},
		{"must set to error", "error", logrus.ErrorLevel},
		{"must not fail and fallback", "random", logrus.InfoLevel},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetLogLevel(tt.lvl)
			assert.Equal(t, logrus.GetLevel(), tt.expected)
		})
	}
}
