package lambdas

import (
	"net/http"

	"github.com/komfy/api/pkg/auth"
)

// RegisterHandler handle the /reg endpoint
func RegisterHandler(resp http.ResponseWriter, req *http.Request) {
	// We check if the header x-www-form-urlencoded exists
	_, ok := req.Header["x-www-form-urlencoded"]
	// If so, the ok variable will be equals to true
	if ok {
		// We then parse the query from url and form
		// Into the variable req.Form and req.PostForm
		req.ParseForm()
		// We create the user based on the PostForm variable
		auth.CreateNewUserWithForm(req.PostForm)

	} else {
		// Else if the header doesn't exist
		// We create the user using the request body
		// Which will be a json object 
		auth.CreateNewUserWithJSON(req.Body)

	}
}
