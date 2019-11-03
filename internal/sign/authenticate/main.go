package authenticate

import (
	"net/http"
	"strings"

	err "github.com/komfy/api/internal/error"
	"github.com/komfy/api/internal/sign"
)

const (
	jSON       string = "application/json"
	urlencoded string = "application/x-www-form-urlencoded"
	formData   string = "multipart/form-data"
)

func User(req *http.Request) sign.Transport {
	if req.Method != http.MethodPost {
		return sign.CreateErrorTransport(err.ErrMethodNotValid)
	}

	content, ok := req.Header["Content-Type"]
	if !ok {
		return sign.CreateErrorTransport(err.ErrContentTypeMissing)
	}

	var userTransport sign.Transport
	if content[0] == jSON {
		userTransport = handleJSON(req)

	} else if content[0] == urlencoded {
		pErr := req.ParseForm()
		if pErr != nil {
			return sign.CreateErrorTransport(pErr)
		}

		userTransport = handleURLEncoded(req.PostForm)

	} else if content = strings.Split(content[0], ";"); content[0] == formData {
		pErr := req.ParseMultipartForm(0)
		if pErr != nil {
			return sign.CreateErrorTransport(pErr)
		}

		userTransport = handleFormData(req.MultipartForm.Value)

	}

	checkChan := make(chan sign.Transport)
	go checkUser(userTransport.User, checkChan)

	jwtTransport := createJWT(userTransport.User)
	if jwtTransport.Error != nil {
		return jwtTransport
	}

	result := <-checkChan
	close(checkChan)
	if result.Error != nil {
		return result
	}

	return jwtTransport
}
