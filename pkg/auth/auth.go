package auth

import (
	"encoding/json"
	"io"
	"net/url"

	"github.com/komfy/api/pkg/jwt"

	"golang.org/x/crypto/bcrypt"

	db "github.com/komfy/api/pkg/database"
	err "github.com/komfy/api/pkg/error"
)

// AuthenticateWithForm is used in order to
// authenticate user based on a urlencoded form
func AuthenticateWithForm(values url.Values) (string, error) {
	username, userExist := values["username"]
	password, passExist := values["password"]

	if !(userExist && passExist) {
		return "", err.ErrValueMissing

	}

	dbUser := db.AskUserByUsername(username[0])

	compareError := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(password[0]))
	if compareError != nil {
		return "", compareError

	}

	token := jwt.CreateToken(dbUser)

	return token, nil
}

// AuthenticateWithJSON is used in order to
// authenticate user based on a json object
func AuthenticateWithJSON(jsonBody io.ReadCloser) (string, error) {
	// Create an empty User object
	user := db.User{}
	// Decode the json object into the user
	json.NewDecoder(jsonBody).Decode(&user)
	dbUser := db.AskUserByUsername(user.Username)

	// If compareError == nil
	// Then both password are the same
	compareError := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	// Delete the password from memory
	user.Password = ""

	if compareError != nil {
		return "", compareError

	}

	token := jwt.CreateToken(dbUser)

	return token, nil
}
