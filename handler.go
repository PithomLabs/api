package handler

import (
	"net/http"

	"github.com/komfy/api/lambdas"
	"github.com/komfy/api/pkg/captcha"
	nu "github.com/komfy/api/pkg/netutils"
)

// MainHandler works as a ServerMux, just in a simpler way
func MainHandler(resp http.ResponseWriter, req *http.Request) {
	// Enable Cross-Origin
	nu.EnableCORS(&resp)

	if !captcha.IsInitialize {
		captcha.InitializeMemoryStorage()

	}

	path := req.URL.Path

	switch path {
	case "/":
		lambdas.IndexHandler(resp, req)

	case "/rand":
		lambdas.RandHandler(resp, req)

	case "/rand_dict":
		lambdas.RandDictHandler(resp, req)

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
