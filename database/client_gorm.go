package database

import (
	"github.com/jclebreton/opensirene/conf"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Support for PostgreSQL dialect
)

// InitGORMClient initializes the main client using the configuration of the
// application
func NewGORMClient() (*gorm.DB, error) {
	return gorm.Open("postgres", conf.C.Database.ConnectionString())
}
