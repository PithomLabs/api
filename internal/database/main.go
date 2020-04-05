package database

import (
	"os"

	// Postgresql driver
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var openDatabase KomfyDB

func InitializeDatabaseInstance(isDev bool) error {
	tempDB, oErr := open(isDev)
	if oErr != nil {
		return oErr
	}

	openDatabase = tempDB
	return nil
}

func open(isDev bool) (KomfyDB, error) {
	dbURL := os.Getenv("DATABASE_URL")
	if isDev {
		dbURL += "?sslmode=disable"
	}

	db, dbErr := gorm.Open("postgres", dbURL)
	if dbErr != nil {
		return KomfyDB{}, dbErr
	}

	db.DB().SetMaxOpenConns(1)

	return KomfyDB{
		Instance: db,
	}, nil
}

func (db KomfyDB) Close() error {
	return db.Instance.Close()
}
