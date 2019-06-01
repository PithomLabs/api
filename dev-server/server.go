package main

import (
	"github.com/komfy/api/lambdas/rand"
	"net/http"
)

func main() {
	http.HandleFunc("/rand", lambdas.rand.Handler)
	http.ListenAndServe(":8080", nil)
}
