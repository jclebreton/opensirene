package conf

import "github.com/sirupsen/logrus"

type Logger struct {
	Level  string `yaml:"level" env:"LOGLEVEL" default:"info"`
	Format string `yaml:"format" env:"LOGFORMAT" default:"text"`
}

func ConfigureLogger(c Logger) {
	SetLogLevel(c.Level)
	SetFormatter(c.Format)
}

// SetLogLevel sets the logging level when possible, otherwise it fallbacks to
// the default logrus level and logs a warning
func SetLogLevel(lvl string) {
	l, err := logrus.ParseLevel(lvl)
	if err != nil {
		logrus.WithField("provided", lvl).Warn("Invalid log level, fallback to Info level")
		logrus.SetLevel(logrus.InfoLevel)
	} else {
		logrus.SetLevel(l)
	}
}

func SetFormatter(format string) {
	switch format {
	case "json":
		logrus.SetFormatter(&logrus.JSONFormatter{})
	default:
	case "text":
		logrus.SetFormatter(&logrus.TextFormatter{})
	}
}
