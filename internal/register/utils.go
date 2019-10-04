package register

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"

	bc "golang.org/x/crypto/bcrypt"

	"github.com/komfy/api/internal/captcha"
	"github.com/komfy/api/internal/database"
	err "github.com/komfy/api/internal/error"
	"github.com/komfy/api/internal/structs"
)

// A simple struct in order to transport information
// from functions
type transport struct {
	User  *structs.User
	Error error
	Bool  bool
}

const (
	// All the different Content-Type header we are accepting
	jSON       string = "application/json"
	urlencoded string = "application/x-www-form-urlencoded"
	multipart  string = "multipart/form-data"

	// Default password cost
	passwordCost int = 8
)

// This function will verify a second time
// if the captcha infos are correct
func doubleCheck(request *http.Request) error {
	capInfos, ok := request.Header["X-Captcha"]
	if !ok {
		return err.ErrCaptchaHeaderMissing
	}

	validInfos := captcha.DoubleCheck(capInfos[0])
	if !validInfos {
		return err.ErrDoubleCheck
	}

	return nil
}

func extractUser(request *http.Request, userChan chan<- transport) {
	content, ok := request.Header["Content-Type"]
	if !ok {
		userChan <- createErrorTransport(err.ErrContentTypeMissing)
		return
	}

	if content[0] == urlencoded {
		pErr := request.ParseForm()
		if pErr != nil {
			userChan <- createErrorTransport(pErr)
			return
		}

		userChan <- parseUrlencoded(request.PostForm)

	} else if content[0] == jSON {
		userChan <- parseJSON(request.Body)

	} else if content = strings.Split(content[0], ";"); content[0] == "multipart/form-data" {
		pErr := request.ParseMultipartForm(0)
		if pErr != nil {
			userChan <- createErrorTransport(pErr)
			return
		}

		userChan <- parseMultipart(request.MultipartForm.Value)

	}
}

func parseUrlencoded(values url.Values) transport {
	username, uExists := values["username"]
	password, pExists := values["password"]
	email, eExists := values["email"]

	if !(uExists && pExists && eExists) {
		return credentialMissing(uExists, pExists, eExists)

	}

	user := &structs.User{
		Username: username[0],
		Password: password[0],
		Email:    email[0],
	}

	return createUserTransport(user)

}

func parseJSON(body io.ReadCloser) transport {
	var user *structs.User
	dErr := json.NewDecoder(body).Decode(user)
	if dErr != nil {
		return createErrorTransport(dErr)

	}

	uExists := user.Username != ""
	pExists := user.Password != ""
	eExists := user.Email != ""
	if !(uExists && pExists && eExists) {
		return credentialMissing(uExists, pExists, eExists)

	}

	return createUserTransport(user)
}

func parseMultipart(values map[string][]string) transport {
	var urlValues url.Values
	for key, value := range values {
		if len(value) == 1 {
			urlValues.Set(key, value[0])
		}
	}

	return parseUrlencoded(urlValues)
}

func isValidUser(db database.KomfyDB, user *structs.User, validChan chan<- transport) {
	valid := db.IsValid(user)
	if !valid {
		validChan <- createErrorTransport(err.ErrUserNotValid)
	}
	validChan <- createBoolTransport(valid)

}

func hashPassword(pass string) (string, error) {
	bytePass, hErr := bc.GenerateFromPassword([]byte(pass), passwordCost)
	if hErr != nil {
		return "", err.ErrHashing
	}

	return string(bytePass), nil
}

func credentialMissing(user, pass, email bool) transport {
	var str string
	if !user {
		str += "username,"
	}
	if !pass {
		str += "password,"
	}
	if !email {
		str += "email,"
	}

	customMessage := "those credentials are missing: " + str[:len(str)-1]
	customError := err.CreateErrorFromString(customMessage)

	return createErrorTransport(customError)
}

// This function create a transport object which have a nil user
// and an error defined by transportError
func createErrorTransport(transportError error) transport {
	return transport{
		User:  nil,
		Error: transportError,
	}
}

func createUserTransport(transportUser *structs.User) transport {
	return transport{
		User:  transportUser,
		Error: nil,
	}
}

func createBoolTransport(transportBool bool) transport {
	return transport{
		Bool:  transportBool,
		Error: nil,
	}
}
