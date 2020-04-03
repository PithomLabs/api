package netutils

import (
	"os"
	"strings"

	"net/http"
)

// IsDev checks if the application runs in a development environment
func IsDev() bool {
	// find out if the server needs to be run in development mode
	isDev := false
	env := os.Getenv("APP_ENV")

	if strings.Contains(env, "dev") {
		isDev = true
	}

	return isDev
}

// EnableCORS add an header to the ResponseWriter
// in order to allow the current handler to receive
// Cross-Origin-Resources-Sharing requests
func EnableCORS(resp *http.ResponseWriter, crossOriginURL string) {
	(*resp).Header().Set("Access-Control-Allow-Origin", crossOriginURL)
	(*resp).Header().Set("Access-Control-Allow-Headers", "Origin, Accept, Content-Type, X-Requested-With, Authorization")
	(*resp).Header().Set("Access-Control-Allow-Methods", "GET,POST")
}
