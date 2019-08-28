package lambdas

import (
	"log"
	"net/http"
	"strings"

	"github.com/komfy/api/pkg/auth"
	nu "github.com/komfy/api/pkg/netutils"
)

const redirectRegURL = "https://komfy.now.sh/verify_email"

// RegisterHandler handle the /reg endpoint
func RegisterHandler(resp http.ResponseWriter, req *http.Request) {
	nu.EnableCORS(&resp)
	if req.Method == http.MethodPost {
		// We collect the header Content-Type
		content, ok := req.Header["Content-Type"]
		// Based on the content-type header value,
		// Use different type of registration
		if !ok {
			resp.WriteHeader(http.StatusBadRequest)
			resp.Write([]byte("You forgot to add a Content-Type header :D"))
			log.Println("Missing content-type header")
			return

		} else if content[0] == "application/x-www-form-urlencoded" {
			// We then parse the query from url and form
			// Into the variable req.Form and req.PostForm
			req.ParseForm()
			// We create the user based on the PostForm variable
			err := auth.CreateNewUserWithForm(resp, req.PostForm)
			if err != nil {
				log.Println(err)
				return
			}

		} else if content[0] == "application/json" {
			// Else if the header doesn't exist
			// We create the user using the request body
			// Which will be a json object
			err := auth.CreateNewUserWithJSON(resp, req.Body)
			if err != nil {
				log.Println(err)
				return
			}

		} else if content = strings.Split(content[0], ";"); content[0] == "multipart/form-data" {
			ferr := req.ParseMultipartForm(0)
			if ferr != nil {
				resp.WriteHeader(http.StatusBadRequest)
				resp.Write([]byte("The multipart/form-data doesn't have a boundary"))
				log.Println(ferr)
				return
			}

			formData := req.MultipartForm.Value

			err := auth.CreateNewUserWithFormData(resp, formData)
			if err != nil {
				log.Println(err)
				return
			}
		}
		// Redirect to komfy main page
		http.Redirect(resp, req, redirectRegURL, http.StatusSeeOther)

	} else {
		resp.WriteHeader(http.StatusMethodNotAllowed)
		resp.Write([]byte("Bad request method"))

	}
}
