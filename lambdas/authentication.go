package lambdas

import (
	"log"
	"net/http"

	"github.com/komfy/api/pkg/auth"
)

const redirectAuthURL = "https://komfy.now.sh"

// AuthenticationHandler handle the /auth endpoint
func AuthenticationHandler(resp http.ResponseWriter, req *http.Request) {
	// We only accept POST request
	if req.Method == http.MethodPost {
		// Create the jwt-token variable
		var jwtToken string
		// Check what type of request it is
		ok := req.Header["Content-Type"][0] == "application/x-www-form-urlencoded"
		// Based on the content-type header,
		// We are going to parse the information
		// Using different functions
		if ok {
			// Parse the form values
			req.ParseForm()

			formToken, err := auth.AuthenticateWithForm(req.PostForm)
			if err != nil {
				resp.WriteHeader(http.StatusBadRequest)
				log.Fatal(err)
				return
			}

			// Update the value of the token
			jwtToken = formToken

		} else {
			// Use the post form body,
			// Which should be a json object
			jsonToken, err := auth.AuthenticateWithJSON(req.Body)
			if err != nil {
				resp.WriteHeader(http.StatusBadRequest)
				log.Fatal(err)
				return
			}
			// Same as line 35
			jwtToken = jsonToken

		}

		// If we reach this place, this mean
		// The authentication went successfully
		cookie := http.Cookie{
			Name:  "jwt-token",
			Value: jwtToken,
		}
		// Add the cookie to the request
		http.SetCookie(resp, &cookie)
		// Redirect after authentication
		http.Redirect(resp, req, redirectAuthURL, http.StatusSeeOther)

	} else {
		resp.WriteHeader(http.StatusMethodNotAllowed)
		log.Fatal("Bad Method")

	}
}
