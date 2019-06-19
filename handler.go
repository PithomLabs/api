package handler

import (
	"fmt"
	"net/http"

	"github.com/komfy/api/lambdas"
)

// MainHandler works as a ServerMux, just in a simpler way
func MainHandler(writer http.ResponseWriter, req *http.Request) {
	path := req.URL.Path

	switch path {
	case "/":
		lambdas.IndexHandler(writer, req)

	case "/rand":
		lambdas.RandHandler(writer, req)

	case "/rand_dict":
		lambdas.RandDictHandler(writer, req)

	case "/dbtest":
		lambdas.DatabaseTest(writer, req)

	default:
		fmt.Fprintf(writer, "Error: Unknown path %s", path)

	}

}
