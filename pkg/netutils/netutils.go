package netutils

import "net/http"

// EnableCORS add an header to the ResponseWriter
// in order to allow the current handler to receive
// Cross-Origin-Ressources-Sharing
func EnableCORS(resp *http.ResponseWriter) {
	(*resp).Header().Set("Access-Control-Allow-Origin", "https://komfy.now.sh/")

}
