package database

import (
	// Golang packages
	"log"
	"os"

	// Community packages
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB is a struct containing the gorm database
type DB struct {
	GDB *gorm.DB
}

// OpenDatabase returns a DB object containing gorm database
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

// CloseDB close the gorm database
func (db *DB) CloseDB() error {
	return db.GDB.Close()
}

// AskUserByID returns the user corresponding to the given ID
func AskUserByID(id string) *User {
	db := OpenDatabase()

	user := &User{}
	db.GDB.First(&user, "user_id = ?", id)

	db.CloseDB()
	return user
}

// IsUserValid check if a user is valid and not a duplicate
func IsUserValid(user *User) bool {
	return false
}

// AddUserToDB add a given user to gorm opened database
func AddUserToDB(user *User) {
	db := OpenDatabase()

	db.GDB.Create(&user)

	db.CloseDB()
}
