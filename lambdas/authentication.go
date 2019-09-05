package lambdas

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/komfy/api/pkg/auth"
	err "github.com/komfy/api/pkg/error"
)

const redirectAuthURL = "https://komfy.now.sh/set_cookie?token=%s"

// AuthenticationHandler handle the /auth endpoint
func AuthenticationHandler(resp http.ResponseWriter, req *http.Request) {
	// We only accept POST request
	if req.Method == http.MethodPost {
		// Create the jwt-token variable
		var jwtToken string
		// Check what type of request it is
		content, ok := req.Header["Content-Type"]
		// Based on the content-type header,
		// We are going to parse the information
		// Using different functions
		if !ok {
			err.HandleErrorInHTTP(resp, err.ErrContentTypeMissing)
			log.Println("Content-Type header is missing")
			return

		} else if content[0] == "application/x-www-form-urlencoded" {
			// Parse the form values
			req.ParseForm()

			formToken, authErr := auth.AuthenticateWithForm(resp, req.PostForm)
			if authErr != nil {
				err.HandleErrorInHTTP(resp, authErr)
				log.Println(authErr)
				return
			}

			// Update the value of the token
			jwtToken = formToken

		} else if content[0] == "application/json" {
			// Use the post form body,
			// Which should be a json object
			jsonToken, authErr := auth.AuthenticateWithJSON(resp, req.Body)
			if authErr != nil {
				err.HandleErrorInHTTP(resp, authErr)
				log.Println(authErr)
				return

			}
			// Same as line 35
			jwtToken = jsonToken

		} else if content = strings.Split(content[0], ";"); content[0] == "multipart/form-data" {
			ferr := req.ParseMultipartForm(0)
			if ferr != nil {
				err.HandleErrorInHTTP(resp, err.ErrMultipartFormData)
				log.Println(ferr)
				return

			}

			formData := req.Form

			dataToken, authErr := auth.AuthenticateWithFormData(resp, formData)
			if authErr != nil {
				err.HandleErrorInHTTP(resp, authErr)
				log.Println(authErr)
				return

			}

			jwtToken = dataToken

		}

		realTokenURL := fmt.Sprintf(redirectAuthURL, jwtToken)

		http.Redirect(resp, req, realTokenURL, http.StatusSeeOther)

	} else {
		resp.WriteHeader(http.StatusMethodNotAllowed)
		resp.Write([]byte("Bad request method"))

	}
}
