package auth

import (
	"encoding/json"
	"fmt"
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
	// Check if we have an username
	username, userExist := formValues["username"]
	// Check if we have an email
	email, emailExist := formValues["email"]
	// Check if we have a password
	password, passExist := formValues["password"]

	// If either the password or the username is missing
	// Returns an error
	if valueMissing := !(passExist && userExist && emailExist); valueMissing {
		var errorMessage string

		if !(userExist && username[0] != "") {
			errorMessage = fmt.Sprintf(err.ErrValueMissingTemplate, "username")

		} else if !(emailExist && email[0] != "") {
			errorMessage = fmt.Sprintf(err.ErrValueMissingTemplate, "email")

		} else if !(passExist && password[0] == "") {
			errorMessage = fmt.Sprintf(err.ErrValueMissingTemplate, "password")

		}

		err.HandleErrorInHTTP(resp, err.CreateError(errorMessage))
		return err.ErrValueMissing
	}

	// Hash the password using bcrypt hash method
	hashedPass, hashError := bc.GenerateFromPassword([]byte(password[0]), passwordCreationCost)
	if hashError != nil {
		err.HandleErrorInHTTP(resp, err.ErrHashing)
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

	return verifyUserAndSendMail(user)

}

// CreateNewUserWithJSON creates a new user based on a json object
func CreateNewUserWithJSON(resp http.ResponseWriter, requestBody io.ReadCloser) error {
	// Create an empty user
	user := &db.User{}
	// Decode the request body and fill the user object with the infos inside
	json.NewDecoder(requestBody).Decode(&user)

	// Check if all credentials exist
	if valueMissing := !(user.Username != "" && user.Email != "" && user.Password != ""); valueMissing {
		var errorMessage string

		if user.Username == "" {
			errorMessage = fmt.Sprintf(err.ErrValueMissingTemplate, "username")

		} else if user.Email == "" {
			errorMessage = fmt.Sprintf(err.ErrValueMissingTemplate, "email")

		} else if user.Password == "" {
			errorMessage = fmt.Sprintf(err.ErrValueMissingTemplate, "password")

		}
		err.HandleErrorInHTTP(resp, err.CreateError(errorMessage))
		return err.ErrValueMissing
	}

	// Hash the user password
	hashedPassword, errCrypt := bc.GenerateFromPassword([]byte(user.Password), passwordCreationCost)
	if errCrypt != nil {
		err.HandleErrorInHTTP(resp, err.ErrHashing)
		return errCrypt

	}

	user.Password = string(hashedPassword)
	hashedPassword = []byte("")

	return verifyUserAndSendMail(user)

}

// VerifyUserAndSendMail is used in order to verify that
// the user is valid and if so,
// send a mail to the given user's email
func verifyUserAndSendMail(user *db.User) error {
	database := db.OpenDatabase()
	defer database.CloseDB()

	if database.IsUserValid(user) {
		database.AddUserToDB(user)
		mail.SendMail(user)

		return nil

	}
	return err.ErrUserNotValid
}
