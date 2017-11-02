package database

import (
	"github.com/jclebreton/opensirene/conf"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Support for PostgreSQL dialect
)

// DB is the main exported client
var DB *gorm.DB

// InitQueryClient initializes the main client using the configuration of the
// application
func InitQueryClient() error {
	var err error
	DB, err = gorm.Open("postgres", conf.C.Database.ConnectionString())
	return err
}
