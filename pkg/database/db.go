package database

import (
	// Golang packages
	"log"
	"os"

	// Community packages
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

type DB struct {
	GDB *gorm.DB
}

func OpenDatabase() *DB {
	if envError := godotenv.Load(".env"); envError != nil {
		log.Fatal(envError)
		return nil

	}

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
	db.GDB.First(&user, "userID = ?", id)

	return user
}
