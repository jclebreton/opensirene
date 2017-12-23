package views

import "github.com/jinzhu/gorm"

type ViewsContext struct {
	GormClient *gorm.DB
	Version    string
	BuildDate  string
}
