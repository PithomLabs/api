package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/komfy/api/pkg/jwt"

	"golang.org/x/crypto/bcrypt"

	db "github.com/komfy/api/pkg/database"
	err "github.com/komfy/api/pkg/error"
)

// AuthenticateWithFormData is used in order to
// authenticate user based on the form-data values
func AuthenticateWithFormData(resp http.ResponseWriter, formValues map[string][]string) (string, error) {
	urlValuesObject := url.Values{}

	for key, value := range formValues {
		if ok := len(value) == 1; ok {
			urlValuesObject.Set(key, value[0])

		}
	}

	return AuthenticateWithForm(resp, urlValuesObject)
}

// AuthenticateWithForm is used in order to
// authenticate user based on a urlencoded form
func AuthenticateWithForm(resp http.ResponseWriter, values url.Values) (string, error) {
	username, userExist := values["username"]
	password, passExist := values["password"]

	// Check for credentials
	if valueMissing := !(passExist && userExist); valueMissing {
		var errorMessage string

		if !(passExist && password[0] != "") {
			errorMessage = fmt.Sprintf(err.ErrValueMissingTemplate, "password")

		} else if !(userExist && username[0] != "") {
			errorMessage = fmt.Sprintf(err.ErrValueMissingTemplate, "username")

		}
		return "", err.CreateError(errorMessage)
	}

	dbUser := db.AskUserByUsername(username[0])

	if !dbUser.Checked {
		return "", err.ErrUserIsntCheck
	}

	compareError := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(password[0]))
	if compareError != nil {
		log.Print(compareError)
		return "", err.ErrBadPassword

	}

	token := jwt.CreateToken(dbUser)

	return token, nil
}

// AuthenticateWithJSON is used in order to
// authenticate user based on a json object
func AuthenticateWithJSON(resp http.ResponseWriter, jsonBody io.ReadCloser) (string, error) {
	// Create an empty User object
	user := db.User{}
	// Decode the json object into the user
	json.NewDecoder(jsonBody).Decode(&user)

	// Check for credentials
	if valueMissing := !(user.Username != "" && user.Email != "" && user.Password != ""); valueMissing {
		var errorMessage string

		if user.Username == "" {
			errorMessage = fmt.Sprintf(err.ErrValueMissingTemplate, "username")

		} else if user.Password == "" {
			errorMessage = fmt.Sprintf(err.ErrValueMissingTemplate, "password")

		}
		err.HandleErrorInHTTP(resp, err.CreateError(errorMessage))
		return "", err.ErrValueMissing
	}

	dbUser := db.AskUserByUsername(user.Username)

	if !dbUser.Checked {
		return "", err.ErrUserIsntCheck
	}

	// If compareError == nil
	// Then both password are the same
	compareError := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	// Delete the password from memory
	user.Password = ""

	if compareError != nil {
		log.Print(compareError)
		return "", err.ErrBadPassword

	}

	token := jwt.CreateToken(dbUser)

	return token, nil
}
