package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/komfy/api/lambdas"
)

func main() {
	fmt.Println("Server is running on port 8080...")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", lambdas.IndexHandler)
	http.HandleFunc("/rand", lambdas.RandHandler)
	http.HandleFunc("/rand_dict", lambdas.RandDictHandler)
	http.HandleFunc("/reg", lambdas.RegisterHandler)
	http.HandleFunc("/graphql", lambdas.GraphQLHandler)
	http.ListenAndServe(":8080", nil)
}
