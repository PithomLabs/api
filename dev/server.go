package main

import (
	h "github.com/komfy/api"
	"net/http"
)

func main() {
	http.HandleFunc("/", h.Handler)
	http.ListenAndServe(":8080", nil)
}
