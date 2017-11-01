package database

import (
	"github.com/jclebreton/opensirene/conf"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func InitQueryClient() error {
	var err error
	DB, err = gorm.Open("postgres", conf.C.Database.ConnectionString())
	return err
}
