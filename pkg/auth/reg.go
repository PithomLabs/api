package auth

import (
	"encoding/json"
	"io"
	"log"
	"net/url"

	db "github.com/komfy/api/pkg/database"
	bc "golang.org/x/crypto/bcrypt"
)

const passwordCreationCost = 8

// CreateNewUserWithForm creates a new user based on the form urlencoded values
func CreateNewUserWithForm(formValues url.Values) {
	// Check if we have a password
	pass, passExists := formValues["password"]
	// Check if we have an username
	username, nameExists := formValues["username"]
	// Check if we have an email
	email, emailExists := formValues["email"]

	// If either the password or the username is missing
	// Returns an error
	if !(passExists && nameExists) {
		log.Fatal("An error occured, values are missing")
	}

	// Hash the password using bcrypt hash method
	hashedPass := bc.GenerateFromPassword([]byte(pass[0]), passwordCreationCost)

	// Create the user and fill it with username and password
	user := &db.User{
		Name:     username[0],
		Password: hashedPass,
	}

	// If the email exists, add it to the user object
	if emailExists {
		user.Email = email[0]
	}

	db.AddUserToDB(user)
}

// CreateNewUserWithJSON creates a new user based on a json object
func CreateNewUserWithJSON(requestBody io.ReadCloser) {
	// Create an empty user
	user := &db.User{}
	// Decode the request body and fill the user object with the infos inside
	json.NewEncoder(requestBody).Decode(&user)
	// Hash the user password
	user.Password = bc.GenerateFromPassword([]byte(user.Password), passwordCreationCost)

	db.AddUserToDB(user)
}
