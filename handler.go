package handler

import (
	"log"
	"net/http"

	"github.com/komfy/api/internal/initialize"
	"github.com/komfy/api/lambdas"

	"github.com/komfy/api/internal/netutils"
)

// MainHandler works as a ServerMux, just in a much simpler way
func MainHandler(resp http.ResponseWriter, req *http.Request) {
	// Enable Cross-Origin
	netutils.EnableCORS(&resp, "https://komfy.now.sh")

	if !initialize.IsOkay {
		iErr := initialize.TurnOkay()
		if iErr != nil {
			log.Println(iErr)
		}

	}

	path := req.URL.Path

	switch path {
	case "/":
		lambdas.IndexHandler(resp, req)

	case "/rand":
		lambdas.PasswordCharacterHandler(resp, req)

	case "/rand_dict":
		lambdas.PasswordDictionnaryHandler(resp, req)

	case "/reg":
		lambdas.RegisterHandler(resp, req)

	case "/auth":
		lambdas.AuthenticationHandler(resp, req)

	case "/verify":
		lambdas.VerifyHandler(resp, req)

	case "/graphql":
		lambdas.GraphQLHandler(resp, req)

	case "/captcha/get":
		lambdas.GetCaptchaHandler(resp, req)

	case "/captcha/verify":
		lambdas.VerifyCaptchaHandler(resp, req)

	default:
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte("Error 404: Unknown path"))

	}

}
