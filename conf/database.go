package conf

import (
	"fmt"
)

// Database is the structure that holds all the mandatory information
// about the database connection
type Database struct {
	User     string `yaml:"user" env:"DB_USER"`
	Name     string `yaml:"name" env:"DB_NAME" default:"opensirene"`
	Password string `yaml:"password" env:"DB_PASSWORD"`
	Host     string `yaml:"host" env:"DB_HOST" default:"127.0.0.1"`
	Port     int    `yaml:"port" env:"DB_PORT" default:"5432"`
	SSLMode  string `yaml:"sslmode" env:"DB_SSL_MODE"`
}

// ConnectionString generates and returns the connection string used by the
// standard SQL library or an ORM like Gorm.
func (d Database) ConnectionString() string {
	return fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%d sslmode=%s",
		d.User,
		d.Password,
		d.Name,
		d.Host,
		d.Port,
		d.SSLMode,
	)
}
