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
	"github.com/komfy/api/internal/mail"
	"github.com/komfy/api/internal/sign"
	"github.com/komfy/api/internal/structs"
)

const (
	// All the different Content-Type header we are accepting
	jSON       string = "application/json"
	urlencoded string = "application/x-www-form-urlencoded"
	multipart  string = "multipart/form-data"

	// Default password cost
	passwordCost int = 8
)

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

func extractUser(request *http.Request, userChan chan<- sign.Transport) {
	content, ok := request.Header["Content-Type"]
	if !ok {
		userChan <- sign.CreateErrorTransport(err.ErrContentTypeMissing)
		return
	}

	if content[0] == urlencoded {
		pErr := request.ParseForm()
		if pErr != nil {
			userChan <- sign.CreateErrorTransport(pErr)
			return
		}

		userChan <- parseUrlencoded(request.PostForm)

	} else if content[0] == jSON {
		userChan <- parseJSON(request.Body)

	} else if content = strings.Split(content[0], ";"); content[0] == "multipart/form-data" {
		pErr := request.ParseMultipartForm(0)
		if pErr != nil {
			userChan <- sign.CreateErrorTransport(pErr)
			return
		}

		userChan <- parseMultipart(request.MultipartForm.Value)

	}
}

func parseUrlencoded(values url.Values) sign.Transport {
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

	return sign.CreateUserTransport(user)

}

func parseJSON(body io.ReadCloser) sign.Transport {
	var user *structs.User
	dErr := json.NewDecoder(body).Decode(&user)
	if dErr != nil {
		return sign.CreateErrorTransport(dErr)

	}

	uExists := user.Username != ""
	pExists := user.Password != ""
	eExists := user.Email != ""
	if !(uExists && pExists && eExists) {
		return credentialMissing(uExists, pExists, eExists)

	}

	return sign.CreateUserTransport(user)
}

func parseMultipart(values map[string][]string) sign.Transport {
	urlValues := url.Values{}
	for key, value := range values {
		if len(value) == 1 {
			urlValues.Set(key, value[0])
		}
	}

	return parseUrlencoded(urlValues)
}

func isValidUser(user *structs.User, validChan chan<- sign.Transport) {
	valid, vErr := database.IsValidUser(user)
	if vErr != nil {
		validChan <- sign.CreateErrorTransport(err.ErrInDatabaseOccured)
	} else if !valid {
		validChan <- sign.CreateErrorTransport(err.ErrUserNotValid)
		return
	}
	validChan <- sign.CreateBoolTransport(valid)

}

func hashPassword(pass string) (string, error) {
	bytePass, hErr := bc.GenerateFromPassword([]byte(pass), passwordCost)
	if hErr != nil {
		return "", err.ErrHashing
	}

	return string(bytePass), nil
}

func credentialMissing(user, pass, email bool) sign.Transport {
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

	return sign.CreateErrorTransport(customError)
}

func sendMail(user *structs.User, sendChan chan sign.Transport) {
	if !mail.IsValid(user.Email) {
		// Receive the empty Transport which is normally supposed
		// to indicate to the mail.Send function
		// when it can access the UserID field
		<-sendChan
		sendChan <- sign.CreateErrorTransport(err.ErrMailNotValid)
	}

	mErr := mail.Send(user, sendChan)
	if mErr != nil {
		sendChan <- sign.CreateErrorTransport(mErr)
	}

	sendChan <- sign.CreateErrorTransport(nil)
}
