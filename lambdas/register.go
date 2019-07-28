package lambdas

import (
	"log"
	"net/http"

	"github.com/komfy/api/pkg/auth"
	nu "github.com/komfy/api/pkg/netutils"
)

const redirectRegURL = "https://komfy.now.sh/"

// RegisterHandler handle the /reg endpoint
func RegisterHandler(resp http.ResponseWriter, req *http.Request) {
	nu.EnableCORS(&resp)
	if req.Method == http.MethodPost {
		// We collect the header Content-Type
		content := req.Header["Content-Type"]
		// We check if the value of this header is
		// application/x-www-form-urlencoded
		ok := content[0] == "application/x-www-form-urlencoded"
		// Based on the ok variable, we use
		// the urlencoded values or the json ones
		if ok {
			// We then parse the query from url and form
			// Into the variable req.Form and req.PostForm
			req.ParseForm()
			// We create the user based on the PostForm variable
			err := auth.CreateNewUserWithForm(req.PostForm)
			if err != nil {
				resp.WriteHeader(http.StatusBadRequest)
				log.Fatal(err)
			}

		} else {
			// Else if the header doesn't exist
			// We create the user using the request body
			// Which will be a json object
			err := auth.CreateNewUserWithJSON(req.Body)
			if err != nil {
				resp.WriteHeader(http.StatusBadRequest)
				log.Fatal(err)
			}

		}

		// Redirect to komfy main page
		http.Redirect(resp, req, redirectRegURL, http.StatusSeeOther)

	} else {
		resp.WriteHeader(http.StatusMethodNotAllowed)

	}
}
