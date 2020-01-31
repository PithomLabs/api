package lambdas

import (
	"net/http"

	"encoding/json"

	"github.com/komfy/api/internal/graphql"
)

// GraphQLHandler handle the /graphql endpoint
func GraphQLHandler(resp http.ResponseWriter, req *http.Request) {
	result := graphql.Do(req)
	// The response will be formatted in json style
	resp.Header().Add("Content-type", "application/json")
	json.NewEncoder(resp).Encode(result)

}
