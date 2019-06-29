package lambdas

import (
	"net/http"

	"encoding/json"

	"github.com/komfy/api/pkg/gql"
)

func GraphQLHandler(writer http.ResponseWriter, req *http.Request) {
	result := gql.Do(req)
	json.NewEncoder(writer).Encode(result)

}
