package database

import (
	"fmt"
	"net/url"
	"os"

	"github.com/jinzhu/gorm"

	// Postgresql driver
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/komfy/api/internal/netutils"
	"github.com/komfy/api/internal/structs"
)

var openDatabase KomfyDB

func InitializeDatabaseInstance() error {
	tempDB, oErr := open()
	if oErr != nil {
		return oErr
	}

	openDatabase = tempDB
	return nil
}

func open() (KomfyDB, error) {
	dbURL, uErr := url.Parse(os.Getenv("DATABASE_URL"))
	if uErr != nil {
		return KomfyDB{}, uErr
	}

	isDev := netutils.IsDev()
	if isDev {
		dbURL.RawQuery = dbURL.RawQuery + "&sslmode=disable"
	}

	db, dbErr := gorm.Open("postgres", dbURL.String())
	if dbErr != nil {
		return KomfyDB{}, dbErr
	}

	db.DB().SetMaxOpenConns(1)

	db.AutoMigrate(&structs.User{}, &structs.Settings{}, &structs.Entity{}, &structs.Asset{}, &structs.Content{})

	return KomfyDB{
		Instance: db,
	}, nil
}

func (db KomfyDB) Close() error {
	return db.Instance.Close()
}
