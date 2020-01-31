package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/komfy/api/internal/initialize"
	"github.com/komfy/api/internal/netutils"
	"github.com/komfy/api/lambdas"
)

func main() {
	fmt.Println("Reading env variables from .env file...")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	if !initialize.IsOkay {
		initialize.TurnOkay()
	}

	fmt.Println("Done...")
	fmt.Println("Server is running on port 8080...")

	http.HandleFunc("/", AddCORSOnLocal)
	http.ListenAndServe(":8080", nil)

}

// AddCORSOnLocal enable CORS request from localhost:3000 when doing local development
func AddCORSOnLocal(resp http.ResponseWriter, req *http.Request) {
	netutils.EnableCORS(&resp, "http://localhost:3000")
	// We suppress the '/' at the beginning of the path
	path := req.URL.Path[1:]

	switch path {
	case "rand":
		lambdas.PasswordCharacterHandler(resp, req)

	case "rand_dict":
		lambdas.PasswordDictionnaryHandler(resp, req)

	case "reg":
		lambdas.RegisterHandler(resp, req)

	case "auth":
		lambdas.AuthenticationHandler(resp, req)

	case "verify":
		lambdas.VerifyHandler(resp, req)

	case "graphql":
		lambdas.GraphQLHandler(resp, req)

	case "captcha/get":
		lambdas.GetCaptchaHandler(resp, req)

	case "captcha/verify":
		lambdas.VerifyCaptchaHandler(resp, req)

	default:
		lambdas.IndexHandler(resp, req)

	}
}
