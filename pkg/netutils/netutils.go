package netutils

import "net/http"

// EnableCORS add an header to the ResponseWriter
// in order to allow the current handler to receive
// Cross-Origin-Ressources-Sharing requests
func EnableCORS(resp *http.ResponseWriter, crossOriginURL string) {
	(*resp).Header().Set("Access-Control-Allow-Origin", crossOriginURL)

}
