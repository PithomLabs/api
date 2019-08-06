package auth

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	db "github.com/komfy/api/pkg/database"
	mail "github.com/komfy/api/pkg/email"
	err "github.com/komfy/api/pkg/error"
	bc "golang.org/x/crypto/bcrypt"
)

const passwordCreationCost = 8

// CreateNewUserWithFormData creates a new url.Values object
// And add the formValue's values to it
// Then call CreateNewUserWithForm
func CreateNewUserWithFormData(resp http.ResponseWriter, formValue map[string][]string) error {
	urlValuesObject := url.Values{}

	for key, value := range formValue {
		if ok := len(value) == 1; ok {
			urlValuesObject.Set(key, value[0])

		}
	}

	return CreateNewUserWithForm(resp, urlValuesObject)
}

// CreateNewUserWithForm creates a new user based on the form urlencoded values
func CreateNewUserWithForm(resp http.ResponseWriter, formValues url.Values) error {
	// Check if we have a password
	pass, passExists := formValues["password"]
	// Check if we have an username
	username, nameExists := formValues["username"]
	// Check if we have an email
	email, emailExists := formValues["email"]

	// If either the password or the username is missing
	// Returns an error
	if !(passExists && nameExists && emailExists) {
		err.HandleErrorInHTTP(resp, err.ErrValueMissing)
		return err.ErrValueMissing
	}

	// Hash the password using bcrypt hash method
	hashedPass, hashError := bc.GenerateFromPassword([]byte(pass[0]), passwordCreationCost)
	if hashError != nil {
		err.HandleErrorInHTTP(resp, "An error occured while hashing")
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
		err.HandleErrorInHTTP(resp, err.ErrUserNotValid)
		return err.ErrUserNotValid

	}

	return nil

}

// CreateNewUserWithJSON creates a new user based on a json object
func CreateNewUserWithJSON(resp http.ResponseWriter, requestBody io.ReadCloser) error {
	// Create an empty user
	user := &db.User{}
	// Decode the request body and fill the user object with the infos inside
	json.NewDecoder(requestBody).Decode(&user)

	if user.Username == "" || user.Email == "" || user.Password == "" {
		err.HandleErrorInHTTP(resp, err.ErrValueMissing)
		return err.ErrValueMissing

	}
	// Hash the user password
	hashedPassword, errCrypt := bc.GenerateFromPassword([]byte(user.Password), passwordCreationCost)
	if errCrypt != nil {
		err.HandleErrorInHTTP(resp, "An error occured while hashing")
		return errCrypt

	}

	user.Password = string(hashedPassword)
	hashedPassword = []byte("")

	if db.IsUserValid(user) {
		db.AddUserToDB(user)
		mail.SendMail(user)

	} else {
		err.HandleErrorInHTTP(resp, err.ErrUserNotValid)
		return err.ErrUserNotValid

	}

	return nil

}
