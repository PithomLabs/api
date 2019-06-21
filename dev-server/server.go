package main

import (
	"net/http"

	"github.com/komfy/api/lambdas"
)

func main() {
	http.HandleFunc("/", lambdas.IndexHandler)
	http.HandleFunc("/rand", lambdas.RandHandler)
	http.HandleFunc("/rand_dict", lambdas.RandDictHandler)
	http.HandleFunc("/dbtests", lambdas.DatabaseTest)
	http.ListenAndServe(":8080", nil)
}
