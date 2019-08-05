package lambdas

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/komfy/api/pkg/auth"
	nu "github.com/komfy/api/pkg/netutils"
)

const redirectAuthURL = "https://komfy.now.sh/set_cookie?token=%s"

// AuthenticationHandler handle the /auth endpoint
func AuthenticationHandler(resp http.ResponseWriter, req *http.Request) {
	nu.EnableCORS(&resp)
	// We only accept POST request
	if req.Method == http.MethodPost {
		// Create the jwt-token variable
		var jwtToken string
		// Check what type of request it is
		content := req.Header["Content-Type"][0]
		// Based on the content-type header,
		// We are going to parse the information
		// Using different functions
		if content == "application/x-www-form-urlencoded" {
			// Parse the form values
			req.ParseForm()

			formToken, err := auth.AuthenticateWithForm(resp, req.PostForm)
			if err != nil {
				resp.WriteHeader(http.StatusBadRequest)
				log.Fatal(err)

			}

			// Update the value of the token
			jwtToken = formToken

		} else if content == "application/json" {
			// Use the post form body,
			// Which should be a json object
			jsonToken, err := auth.AuthenticateWithJSON(resp, req.Body)
			if err != nil {
				resp.WriteHeader(http.StatusBadRequest)
				log.Fatal(err)

			}
			// Same as line 35
			jwtToken = jsonToken

		} else if content = strings.Split(content, ";")[0]; content == "multipart/form-data" {
			req.ParseMultipartForm(0)

			formData := req.Form

			dataToken, err := auth.AuthenticateWithFormData(resp, formData)
			if err != nil {
				resp.WriteHeader(http.StatusBadRequest)
				log.Fatal(err)
			}

			jwtToken = dataToken

		}

		realTokenURL := fmt.Sprintf(redirectAuthURL, jwtToken)

		http.Redirect(resp, req, realTokenURL, http.StatusSeeOther)

	} else {
		// Write an error message when request isn't post
		resp.Write([]byte("Bad request method"))
		resp.WriteHeader(http.StatusMethodNotAllowed)

	}
}
