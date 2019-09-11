package database

import (

	// Golang packages
	"log"
	"os"

	// Community packages
	"github.com/jinzhu/gorm"
	// Driver for postgresql
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB is a struct containing the gorm database
type DB struct {
	GDB *gorm.DB
}

// OpenDatabase returns a DB object containing gorm database
func OpenDatabase() *DB {
	db, dbError := gorm.Open("postgres", os.Getenv("database"))
	if dbError != nil {
		log.Print(dbError)
		return nil

	}

	return &DB{
		GDB: db,
	}

}

// CloseDB close the gorm database
func (db *DB) CloseDB() {
	err := db.GDB.Close()
	if err != nil {
		log.Print("Database as encountered a problem")
	}
}

// AskUserByID returns the user corresponding to the given ID
func (db *DB) AskUserByID(id string) *User {
	user := &User{}
	db.GDB.First(&user, "user_id = ?", id)

	return user
}

// AskUserByUsername return the first user (and only one)
// that correspond to the given username
func (db *DB) AskUserByUsername(username string) *User {
	user := &User{}
	db.GDB.First(&user, "username = ?", username)

	return user

}

// GetSameUsers returns a slice which contain
// All the db users that have a field identical to
// The template user
func (db *DB) GetSameUsers(userTemplate *User) []User {
	users := []User{}

	db.GDB.Where("email = ? OR username = ?", userTemplate.Email, userTemplate.Username).Find(&users)

	return users
}

// IsUserValid check if a user is valid and not a duplicate
func (db *DB) IsUserValid(user *User) bool {
	users := db.GetSameUsers(user)

	// If the slice lengthy is > 0
	// Then there is already a user that have
	// the same email adress or the same username
	return !(len(users) > 0)
}

// AddUserToDB add a given user to gorm opened database
func (db *DB) AddUserToDB(user *User) {
	db.GDB.Create(&user)

}

// UpdateCheckValue change the check value and allow
// the user to authenticate to komfy
func (db *DB) UpdateCheckValue(user *User) {
	db.GDB.Model(&user).Update("Checked", true)

}

// AskPostByID return a post given an id
func (db *DB) AskPostByID(id string) *Post {
	post := &Post{}

	db.GDB.Where("post_id = ?", id).First(&post)

	return post
}
