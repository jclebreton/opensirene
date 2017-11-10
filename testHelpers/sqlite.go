// +build !netgo

// Package testhelpers : since we use this netgo flag when building from script, this file will be exclued from build
// purpose : just to be sure these helpers are just used for tests !
package testHelpers

import (
	"github.com/jinzhu/gorm"

	// sqlite for local tests ...
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// GetConn helper
func GetConn() *gorm.DB {
	db, err := gorm.Open("sqlite3", "test.db") // relative to caller file
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

// Reset : cleans the DB
func Reset(db *gorm.DB, table interface{}) {
	db.DropTable(table)
	db.AutoMigrate(table)
}
