package netutils

import "net/http"

// EnableCORS add an header to the ResponseWriter
// in order to allow the current handler to receive
// Cross-Origin-Resources-Sharing requests
func EnableCORS(resp *http.ResponseWriter, crossOriginURL string) {
	(*resp).Header().Set("Access-Control-Allow-Origin", crossOriginURL)
	(*resp).Header().Set("Access-Control-Allow-Headers", "Origin, Accept, Content-Type, X-Requested-With, Authorization")
	(*resp).Header().Set("Access-Control-Allow-Methods", "GET,POST")
}
