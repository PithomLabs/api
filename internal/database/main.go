package database

import (
	"os"

	"github.com/jinzhu/gorm"

	// Postgresql driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
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
	db, dbErr := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
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
