package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/friendsofgo/graphiql"
	"github.com/joho/godotenv"

	"github.com/komfy/api/internal/initialize"
	"github.com/komfy/api/internal/netutils"
	"github.com/komfy/api/lambdas"
)

var (
	frontendURL = os.Getenv("FRONTEND_URL")
)

const (
	defaultPort string = "8080"
)

// TODO: Comment the whole API

func main() {
	// find out if the server needs to be run in development mode
	isDev := netutils.IsDev()

	// Doing some tweak when working on local server
	if isDev {
		fmt.Println("Running in development mode")

		// Adding GraphiQL handler to the local server
		handler, giErr := graphiql.NewGraphiqlHandler("/graphql")
		if giErr != nil {
			panic(giErr)
		}

		fmt.Printf("  --> You can access GraphiQL at /graphiql\n\n")
		http.Handle("/graphiql", handler)
	}

	// Trying to get rooted path of the current directory we are launching
	// binary from
	cwd, cErr := os.Getwd()
	if cErr != nil {
		fmt.Printf("Could not get the working directory, using the binary location instead\n\n")
		cwd = ""
	}

	// This is the default rooted path: /..../.env
	envPathTemplate := path.Join(cwd, ".env")
	// Those are all the paths we need to check for environment variables
	// We first test the /..../.env.dev file, then the /..../.env file
	envFiles := []string{
		envPathTemplate + ".dev",
		envPathTemplate,
	}

	fmt.Printf("Trying to read env variables...\n")

	var eErr error
	for _, envFile := range envFiles {
		eErr = godotenv.Load(envFile)
		if eErr != nil {
			fmt.Printf("  --> The file %s can't be read... Trying next possibility\n", envFile)
			continue
		}
		fmt.Printf("  --> Environment variables were read from the file %s\n\n", envFile)
		break
	}
	if eErr != nil {
		fmt.Println()
		log.Fatal("No env files were able to be read, please create one following .env.example\n\n")
	}

	// Here we won't log.Fatal because we know that
	// initialize.TurnOkay will give us an error from the Schema.
	// If you have any other errors, you must terminate the currently running
	// server, and check docs or create an issue to solve it
	if !initialize.IsOkay {
		iErrs := initialize.TurnOkay(isDev)
		if iErrs != nil {
			fmt.Println("Errors occured during initialization:")
			for _, iErr := range iErrs {
				fmt.Printf("  --> %s\n", iErr.Error())
			}
			fmt.Println()
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		// Only warn about this if we're in production
		if !isDev {
			fmt.Printf("Could not get the port, falling back to %s instead\n\n", defaultPort)
		}

		port = defaultPort
	}

	fmt.Printf("Server is running on port %s\n", port)

	http.HandleFunc("/", mainHandler)
	http.ListenAndServe(":"+port, nil)

}

// mainHandler (was known as AddCORSOnLocal) handles everything
func mainHandler(resp http.ResponseWriter, req *http.Request) {
	netutils.EnableCORS(&resp, frontendURL)
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
