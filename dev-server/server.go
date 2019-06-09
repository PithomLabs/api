package main

import (
	"fmt"
	"net/http"

	"github.com/komfy/api/lambdas/index"
	"github.com/komfy/api/lambdas/rand"
	"github.com/komfy/api/lambdas/randdict"
	db "github.com/komfy/api/pkg/database"
)

func main() {
	// Little DB test
	db := db.CreateDatabase()
	fmt.Println(db.AskUserOfID("10"))

	http.HandleFunc("/", index.Handler)
	http.HandleFunc("/rand", rand.Handler)
	http.HandleFunc("/randDict", randdict.Handler)
	http.ListenAndServe(":8080", nil)
}
