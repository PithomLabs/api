package authenticate

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/komfy/api/internal/database"
	err "github.com/komfy/api/internal/error"
	"github.com/komfy/api/internal/jwt"
	"github.com/komfy/api/internal/sign"
	"github.com/komfy/api/internal/structs"
	"golang.org/x/crypto/bcrypt"
)

func handleJSON(req *http.Request) sign.Transport {
	var user *structs.User
	dErr := json.NewDecoder(req.Body).Decode(&user)
	if dErr != nil {
		return sign.CreateErrorTransport(dErr)
	}

	uExists := user.Username != ""
	pExists := user.Password != ""

	if !uExists || !pExists {
		return credentialMissing(uExists, pExists)
	}

	return sign.CreateUserTransport(user)
}

func handleURLEncoded(values url.Values) sign.Transport {
	username, uExists := values["username"]
	password, pExists := values["password"]
	if !uExists || !pExists {
		return credentialMissing(uExists, pExists)
	}

	user := &structs.User{
		Username: username[0],
		Password: password[0],
	}

	return sign.CreateUserTransport(user)

}

func handleFormData(multiValues map[string][]string) sign.Transport {
	var values url.Values
	for key, value := range multiValues {
		if len(value) == 1 {
			values.Set(key, value[0])
		}
	}

	return handleURLEncoded(values)
}

func credentialMissing(user, pass bool) sign.Transport {
	var missing bytes.Buffer

	_, wErr := missing.WriteString("those creadentials are missing: ")
	if wErr != nil {
		return sign.CreateErrorTransport(wErr)
	}

	if !user {
		_, wErr := missing.WriteString("username,")
		if wErr != nil {
			return sign.CreateErrorTransport(wErr)
		}

	}
	if !pass {
		_, wErr := missing.WriteString("password,")
		if wErr != nil {
			return sign.CreateErrorTransport(wErr)
		}
	}

	customError := err.CreateErrorFromString(missing.String())

	return sign.CreateErrorTransport(customError)
}

func checkUser(user *structs.User, checkChan chan sign.Transport) {
	dbUser := database.UserByName(user.Username)
	if dbUser.Username == "" {
		checkChan <- sign.CreateErrorTransport(err.ErrUserDoesntExist)
		return
	}

	if !dbUser.Checked {
		checkChan <- sign.CreateErrorTransport(err.ErrUserIsntCheck)
		return
	}

	cErr := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if cErr != nil {
		checkChan <- sign.CreateErrorTransport(err.ErrBadPassword)
		return
	}

	checkChan <- sign.CreateErrorTransport(nil)
}

func createJWT(user *structs.User) sign.Transport {
	jwt, jErr := jwt.Create(user)
	if jErr != nil {
		return sign.CreateErrorTransport(jErr)
	}

	return sign.CreateJWTTransport(jwt)
}
