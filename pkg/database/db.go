package database

import (
	// Golang packages
	"log"
	"os"

	// Community packages
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DB struct {
	GDB *gorm.DB
}

func OpenDatabase() *DB {
	db, dbError := gorm.Open("postgres", os.Getenv("URL"))
	if dbError != nil {
		log.Fatal(dbError)
		return nil

	}

	return &DB{
		GDB: db,
	}

}

func (db *DB) CloseDB() error {
	return db.GDB.Close()
}

func (db *DB) AskUserByID(id string) *User {
	user := &User{}
	db.GDB.First(&user, "user_id = ?", id)

	return user
}
