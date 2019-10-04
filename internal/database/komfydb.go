package database

import "github.com/jinzhu/gorm"

// KomfyDB is the komfy database
type KomfyDB struct {
	Instance *gorm.DB
}
