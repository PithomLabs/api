package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/komfy/api/lambdas"
)

func main() {
	fmt.Println("Reading env variables from .env file...")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Done...")
	fmt.Println("Server is running on port 8080...")
	http.HandleFunc("/", lambdas.IndexHandler)
	http.HandleFunc("/rand", lambdas.RandHandler)
	http.HandleFunc("/rand_dict", lambdas.RandDictHandler)
	http.HandleFunc("/reg", lambdas.RegisterHandler)
	http.HandleFunc("/graphql", lambdas.GraphQLHandler)
	http.ListenAndServe(":8080", nil)

}
