package main

import (
	"net/http"

	"github.com/komfy/api/lambdas"

	//db "github.com/komfy/api/pkg/database"
)

func main() {
	http.HandleFunc("/", lambdas.IndexHandler)
	http.HandleFunc("/rand", lambdas.RandHandler)
	http.HandleFunc("/rand_dict", lambdas.RandDictHandler)
	http.ListenAndServe(":8080", nil)
}
