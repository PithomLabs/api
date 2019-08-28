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
		log.Fatal(dbError)
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
		log.Fatal("Database as encountered a problem")
	}
}

// AskUserByID returns the user corresponding to the given ID
func AskUserByID(id string) *User {
	db := OpenDatabase()
	defer db.CloseDB()

	user := &User{}
	db.GDB.First(&user, "user_id = ?", id)

	return user
}

// AskUserByUsername return the first user (and only one)
// that correspond to the given username
func AskUserByUsername(username string) *User {
	db := OpenDatabase()
	defer db.CloseDB()

	user := &User{}
	db.GDB.First(&user, "username = ?", username)

	return user

}

// GetSameUsers returns a slice which contain
// All the db users that have a field identical to
// The template user
func GetSameUsers(userTemplate *User) []User {
	users := []User{}
	db := OpenDatabase()
	defer db.CloseDB()

	db.GDB.Where("email = ? OR username = ?", userTemplate.Email, userTemplate.Username).Find(&users)

	return users
}

// IsUserValid check if a user is valid and not a duplicate
func IsUserValid(user *User) bool {
	users := GetSameUsers(user)

	// If the slice lengthy is > 0
	// Then there is already a user that have
	// the same email adress or the same username
	return !(len(users) > 0)
}

// AddUserToDB add a given user to gorm opened database
func AddUserToDB(user *User) {
	db := OpenDatabase()
	defer db.CloseDB()

	db.GDB.Create(&user)

}

// UpdateCheckValue change the check value and allow
// the user to authenticate to komfy
func UpdateCheckValue(user *User) {
	db := OpenDatabase()
	defer db.CloseDB()

	db.GDB.Model(&user).Update("Checked", true)
}

// AskPostByID return a post given an id
func AskPostByID(id string) *Post {
	post := &Post{}
	db := OpenDatabase()
	defer db.CloseDB()

	db.GDB.Where("post_id = ?", id).First(&post)

	return post
}
