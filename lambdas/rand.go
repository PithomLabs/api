package lambdas

//Package use for the endpoint /rand and /rand-dict
//Which print a random generated password on the html page
//Which could be used by /auth when creating an account

import (
	"fmt"
	"net/http"

	ru "github.com/komfy/api/pkg/randutils"
)

// RandHandler corresponds to the "/rand" endpoints
func RandHandler(writer http.ResponseWriter, request *http.Request) {
	password := ru.GeneratePassword()
	fmt.Fprintln(writer, password)

}
