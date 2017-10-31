package conf

import (
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
)

// Conf holds the necessary configuration for the application to work
type Conf struct {
	Database     Database   `yaml:"database"`
	Server       Server     `yaml:"server"`
	LogLevel     string     `yaml:"loglevel" env:"LOGLEVEL" default:"info"`
	DownloadPath string     `yaml:"download_path" env:"DOWNLOAD_PATH" default:"downloads"`
	Prometheus   Prometheus `yaml:"prometheus"`
}

// Prometheus is a simple struct holding configuration variables for prometheus
type Prometheus struct {
	Prefix string `yaml:"prefix" env:"PROMETHEUS_PREFIX" default:"opensirene"`
}

// C is the main exported configuration
var C Conf

// Parse will parse every nested fields with the env/defaults parser
// and set the values accordingly
func (c *Conf) Parse() error {
	var err error
	if err = Parse(&c.Database); err != nil {
		return errors.Wrap(err, "couldn't parse Database struct")
	}
	if err = Parse(&c.Server); err != nil {
		return errors.Wrap(err, "couldn't parse Server struct")
	}
	if err = Parse(&c.Server.Cors); err != nil {
		return errors.Wrap(err, "couldn't parse Server.Cors struct")
	}
	if err = Parse(&c.Prometheus); err != nil {
		return errors.Wrap(err, "couldn't parse Prometheus struct")
	}
	if err = Parse(c); err != nil {
		return errors.Wrap(err, "couldn't parse Conf struct")
	}
	switch c.LogLevel {
	case "panic":
		logrus.SetLevel(logrus.PanicLevel)
	case "fatal":
		logrus.SetLevel(logrus.FatalLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	default:
		logrus.WithField("provided", c.LogLevel).Warn("Invalid log level, fallback to Info level")
	}
	if _, err = os.Stat(c.DownloadPath); os.IsNotExist(err) {
		os.MkdirAll(c.DownloadPath, os.ModePerm)
	}
	return Parse(c)
}

// Load loads the configuration file into C
func Load(fp string) error {
	var err error
	var c []byte

	if c, err = ioutil.ReadFile(fp); err != nil {
		return err
	}
	if err = yaml.Unmarshal(c, &C); err != nil {
		return err
	}
	return C.Parse()
}
