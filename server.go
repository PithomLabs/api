package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/friendsofgo/graphiql"
	"github.com/joho/godotenv"

	"github.com/komfy/api/internal/initialize"
	"github.com/komfy/api/internal/netutils"
	"github.com/komfy/api/lambdas"
)

// TODO: Comment the whole API

func main() {

	// find out if the server needs to be run in development mode
	isDev := false
	env := os.Getenv("APP_ENV")

	if strings.Contains(env, "dev") {
		isDev = true
		fmt.Println("Running in development mode")
	}

	// load the .env file, it contains all settings
	envFilePrefix := ""
	if isDev {
		envFilePrefix = ".dev"
	}

	cwd, cErr := os.Getwd()
	if cErr != nil {
		// TODO: come up with a better way to handle this
		panic(cErr)
	}

	envFile := path.Join(cwd, ".env"+envFilePrefix)

	fmt.Printf("Reading env variables from %s...\n", envFile)
	eErr := godotenv.Load(envFile)
	if eErr != nil {
		log.Fatal(eErr)
	}

	if !initialize.IsOkay {
		initialize.TurnOkay()
	}

	fmt.Println("Done...")
	fmt.Println("Server is running on port 8080...")

	if isDev {
		handler, err := graphiql.NewGraphiqlHandler("/graphql")
		if err != nil {
			panic(err)
		}

		http.Handle("/graphiql", handler)

		fmt.Println("You can access GraphiQL at /graphiql")
	}

	http.HandleFunc("/", mainHandler)
	http.ListenAndServe(":8080", nil)

}

// mainHandler (was known as AddCORSOnLocal) handles everything
func mainHandler(resp http.ResponseWriter, req *http.Request) {
	netutils.EnableCORS(&resp, os.Getenv("base_url"))
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
