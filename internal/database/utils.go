package database

import (
	"os"

	"github.com/jinzhu/gorm"
	"github.com/komfy/api/internal/structs"

	// Postgresql driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Open connect to the database, create a
// variable in order to communicate with it, and return
//  a KomfyDB struct plus an error
func Open() (KomfyDB, error) {
	db, dbErr := gorm.Open("postgres", os.Getenv("database"))
	if dbErr != nil {
		return KomfyDB{}, dbErr
	}

	db.DB().SetMaxOpenConns(1)

	return KomfyDB{
		Instance: db,
	}, nil
}

// Close disconnect the Instance inside the db (KomfyDB)
func (db KomfyDB) Close() error {
	return db.Instance.Close()
}

// IsValid tell us if the given user hasn't
// credentials that are already used
func (db KomfyDB) IsValid(user *structs.User) bool {
	users := db.FindUsers(user.Username, user.Email)

	return len(users) == 0
}

// FindUsers returns an array of users with the same
// email or username
func (db KomfyDB) FindUsers(username, email string) []structs.User {
	var users []structs.User
	db.Instance.Where("email = ? OR username = ?", email, username).Find(&users)

	return users
}

func (db KomfyDB) AddUser(user *structs.User) {
	db.Instance.Create(user)
}

func (db KomfyDB) DeleteUser(user *structs.User) {
	db.Instance.Delete(user, "username = ?", user.Username)
}
