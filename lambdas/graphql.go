package lambdas

import (
	"net/http"

	"github.com/komfy/api/internal/graphql"
)

// GraphQLHandler handle the /graphql endpoint
func GraphQLHandler(resp http.ResponseWriter, req *http.Request) {
	result, err := graphql.ExecuteQuery(req)
	if err != nil {
		http.Error(resp, "Couldn't parse json: "+err.Error(), 404)
		return
	}
	// The response will be formatted in json style
	resp.Header().Add("Content-type", "application/json")
	resp.Write(result)
}
