package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/friendsofgo/graphiql"
	"github.com/joho/godotenv"

	"github.com/komfy/api/internal/initialize"
	"github.com/komfy/api/internal/netutils"
	"github.com/komfy/api/lambdas"
)

const defaultPort = 8080

// TODO: Comment the whole API

func main() {

	// find out if the server needs to be run in development mode
	isDev := netutils.IsDev()

	if isDev {
		fmt.Println("Running in development mode")
	}

	// load the .env file, it contains all settings
	envFilePrefix := ""
	if isDev {
		envFilePrefix = ".dev"
	}

	cwd, cErr := os.Getwd()
	if cErr != nil {
		fmt.Println("Could not get the working directory, using the binary location instead")
		cwd = ""
	}

	envFile := path.Join(cwd, ".env"+envFilePrefix)

	fmt.Printf("Reading env variables from %s...\n", envFile)
	eErr := godotenv.Load(envFile)
	if eErr != nil {
		log.Printf("Could not read %s, relying on environment variables instead\n", envFile)
	}

	if !initialize.IsOkay {
		initialize.TurnOkay()
	}

	port, aErr := strconv.Atoi(os.Getenv("PORT"))
	if aErr != nil {
		// only warn about this if we're in production
		if !isDev {
			log.Printf("Could not get the port, falling back to %d instead\n", defaultPort)
		}

		port = defaultPort
	}

	fmt.Printf("Server is running on port %d\n", port)

	if isDev {
		handler, err := graphiql.NewGraphiqlHandler("/graphql")
		if err != nil {
			panic(err)
		}

		http.Handle("/graphiql", handler)

		fmt.Println("You can access GraphiQL at /graphiql")
	}

	http.HandleFunc("/", mainHandler)
	http.ListenAndServe(":"+strconv.Itoa(port), nil)

}

// mainHandler (was known as AddCORSOnLocal) handles everything
func mainHandler(resp http.ResponseWriter, req *http.Request) {
	netutils.EnableCORS(&resp, os.Getenv("FRONTEND_URL"))
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
