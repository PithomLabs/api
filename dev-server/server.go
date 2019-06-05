package main

import (
	index "github.com/komfy/api/lambdas"
	"github.com/komfy/api/lambdas/rand"
	"github.com/komfy/api/lambdas/randDict"
	"net/http"
)

func main() {
	http.HandleFunc("/", index.Handler)
	http.HandleFunc("/rand", rand.Handler)
	http.HandleFunc("/randDict", randDict.Handler)
	http.ListenAndServe(":8080", nil)
}
