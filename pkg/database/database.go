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
	psqlDB *gorm.DB
}

func CreateDatabase() *DB {
	if envError := godotenv.Load("db.env"); envError != nil {
		log.Fatal(envError)
		return nil

	}

	db, dbError := gorm.Open("postgres", os.Getenv("URL"))
	if dbError != nil {
		log.Fatal(dbError)
		return nil

	}

	return &DB{
		psqlDB: db,
	}

}

func (db *DB) AskUserOfID(id string) *User {
	user := &User{}
	db.psqlDB.First(user, id)

	return user
}
