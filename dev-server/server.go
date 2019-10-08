package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/komfy/api/lambdas"
	//"github.com/komfy/api/pkg/captcha"
	nu "github.com/komfy/api/pkg/netutils"
)

func main() {
	fmt.Println("Reading env variables from .env file...")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	/*if !captcha.IsInitialize {
		captcha.InitializeMemoryStorage()
	}*/

	fmt.Println("Done...")
	fmt.Println("Server is running on port 8080...")

	http.HandleFunc("/", AddCORSOnLocal)
	http.HandleFunc("/rand", AddCORSOnLocal)
	http.HandleFunc("/rand_dict", AddCORSOnLocal)
	http.HandleFunc("/reg", AddCORSOnLocal)
	http.HandleFunc("/verify", AddCORSOnLocal)
	http.HandleFunc("/auth", AddCORSOnLocal)
	http.HandleFunc("/captcha/get", AddCORSOnLocal)
	http.HandleFunc("/captcha/verify", AddCORSOnLocal)
	http.HandleFunc("/graphql", AddCORSOnLocal)
	http.ListenAndServe(":8080", nil)

}

// AddCORSOnLocal enable CORS request from localhost:3000 when doing local development
func AddCORSOnLocal(resp http.ResponseWriter, req *http.Request) {
	nu.EnableCORS(&resp, "http://localhost:3000")

	switch req.URL.Path {
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

	}
}
