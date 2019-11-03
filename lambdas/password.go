package lambdas

import (
	"net/http"

	"github.com/komfy/api/internal/password"
)

// PasswordCharacterHandler corresponds to the "/rand" endpoint
func PasswordCharacterHandler(resp http.ResponseWriter, request *http.Request) {
	pass := password.CharacterSequence()
	resp.Write([]byte(pass))

}

// PasswordDictionnaryHandler corresponds to the "/rand_dict" endpoint
func PasswordDictionnaryHandler(resp http.ResponseWriter, request *http.Request) {
	pass := password.WordsSequence()
	resp.Write([]byte(pass))
}
