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
		content, ok := req.Header["Content-Type"]
		// Based on the content-type header,
		// We are going to parse the information
		// Using different functions
		if !ok {
			resp.WriteHeader(http.StatusBadRequest)
			resp.Write([]byte("The multipart/form-data doesn't have a boundary"))
			log.Println("Content-Type header is missing")
			return

		} else if content[0] == "application/x-www-form-urlencoded" {
			// Parse the form values
			req.ParseForm()

			formToken, err := auth.AuthenticateWithForm(resp, req.PostForm)
			if err != nil {
				log.Println(err)
				return
			}

			// Update the value of the token
			jwtToken = formToken

		} else if content[0] == "application/json" {
			// Use the post form body,
			// Which should be a json object
			jsonToken, err := auth.AuthenticateWithJSON(resp, req.Body)
			if err != nil {
				log.Println(err)
				return

			}
			// Same as line 35
			jwtToken = jsonToken

		} else if content = strings.Split(content[0], ";"); content[0] == "multipart/form-data" {
			ferr := req.ParseMultipartForm(0)
			if ferr != nil {
				resp.WriteHeader(http.StatusBadRequest)
				resp.Write([]byte("The multipart/form-data doesn't have a boundary"))
				log.Println(ferr)
				return

			}

			formData := req.Form

			dataToken, err := auth.AuthenticateWithFormData(resp, formData)
			if err != nil {
				log.Println(err)
				return

			}

			jwtToken = dataToken

		}

		realTokenURL := fmt.Sprintf(redirectAuthURL, jwtToken)

		http.Redirect(resp, req, realTokenURL, http.StatusSeeOther)

	} else {
		// Write an error message when request isn't post
		resp.WriteHeader(http.StatusMethodNotAllowed)
		resp.Write([]byte("Bad request method"))

	}
}
