package auth

import (
	"encoding/json"
	"io"
	"log"
	"net/url"

	db "github.com/komfy/api/pkg/database"
	mail "github.com/komfy/api/pkg/email"
	err "github.com/komfy/api/pkg/error"
	bc "golang.org/x/crypto/bcrypt"
)

const passwordCreationCost = 8

// CreateNewUserWithForm creates a new user based on the form urlencoded values
func CreateNewUserWithForm(formValues url.Values) error {
	// Check if we have a password
	pass, passExists := formValues["password"]
	// Check if we have an username
	username, nameExists := formValues["username"]
	// Check if we have an email
	email, emailExists := formValues["email"]

	// If either the password or the username is missing
	// Returns an error
	if !(passExists && nameExists && emailExists) {
		log.Fatal("TEST")
		return err.ErrValueMissing
	}

	// Hash the password using bcrypt hash method
	hashedPass, hashError := bc.GenerateFromPassword([]byte(pass[0]), passwordCreationCost)
	if hashError != nil {
		return hashError
	}

	// Create the user and fill it with username, password and email
	user := &db.User{
		Username: username[0],
		Password: string(hashedPass),
		Email:    email[0],
	}
	// Deleting the password from the hashedPass variable
	hashedPass = []byte("")

	// We are checking if the user isn't a duplicate of another one
	// If not...
	if db.IsUserValid(user) {
		// ...we add this user to database
		// And send him a verification email
		db.AddUserToDB(user)
		mail.SendMail(user)

	} else {
		return err.ErrUserNotValid

	}

	return nil

}

// CreateNewUserWithJSON creates a new user based on a json object
func CreateNewUserWithJSON(requestBody io.ReadCloser) error {
	// Create an empty user
	user := &db.User{}
	// Decode the request body and fill the user object with the infos inside
	json.NewDecoder(requestBody).Decode(&user)

	if user.Username == "" || user.Email == "" {
		return err.ErrValueMissing

	}
	// Hash the user password
	hashedPassword, errCrypt := bc.GenerateFromPassword([]byte(user.Password), passwordCreationCost)
	if errCrypt != nil {
		return errCrypt

	}

	user.Password = string(hashedPassword)
	hashedPassword = []byte("")

	if db.IsUserValid(user) {
		db.AddUserToDB(user)
		mail.SendMail(user)

	} else {
		return err.ErrUserNotValid

	}

	return nil

}
