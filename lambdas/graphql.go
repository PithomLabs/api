package lambdas

import (
	"net/http"

	"encoding/json"

	"github.com/komfy/api/pkg/gql"
)

// GraphQLHandler handle the /graphql endpoint
func GraphQLHandler(writer http.ResponseWriter, req *http.Request) {
	result := gql.Do(req)
	// The response will be formatted in json style
	writer.Header().Add("Content-type", "application/json")
	json.NewEncoder(writer).Encode(result)

}
