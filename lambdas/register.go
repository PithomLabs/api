package lambdas

import (
	"log"
	"net/http"
	"strings"

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
		// Based on the content-type header value,
		// Use different type of registration
		if content[0] == "application/x-www-form-urlencoded" {
			// We then parse the query from url and form
			// Into the variable req.Form and req.PostForm
			req.ParseForm()
			// We create the user based on the PostForm variable
			err := auth.CreateNewUserWithForm(resp, req.PostForm)
			if err != nil {
				resp.WriteHeader(http.StatusBadRequest)
				log.Fatal(err)
			}

		} else if content[0] == "application/json" {
			// Else if the header doesn't exist
			// We create the user using the request body
			// Which will be a json object
			err := auth.CreateNewUserWithJSON(resp, req.Body)
			if err != nil {
				resp.WriteHeader(http.StatusBadRequest)
				log.Fatal(err)
			}

		} else if content = strings.Split(content[0], ";"); content[0] == "multipart/form-data" {
			req.ParseMultipartForm(0)

			formData := req.MultipartForm.Value

			err := auth.CreateNewUserWithFormData(resp, formData)
			if err != nil {
				resp.WriteHeader(http.StatusBadRequest)
				log.Fatal(err)
			}
		}
		// Redirect to komfy main page
		http.Redirect(resp, req, redirectRegURL, http.StatusSeeOther)

	} else {
		resp.Write([]byte("Bad request method"))
		resp.WriteHeader(http.StatusMethodNotAllowed)

	}
}
