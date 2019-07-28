package lambdas

import (
	"net/http"

	"encoding/json"

	"github.com/komfy/api/pkg/gql"
	nu "github.com/komfy/api/pkg/netutils"
)

// GraphQLHandler handle the /graphql endpoint
func GraphQLHandler(resp http.ResponseWriter, req *http.Request) {
	nu.EnableCORS(&resp)
	result := gql.Do(req)
	// The response will be formatted in json style
	resp.Header().Add("Content-type", "application/json")
	json.NewEncoder(resp).Encode(result)

}
