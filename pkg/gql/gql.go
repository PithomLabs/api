package gql

import (
	"context"
	"net/http"

	"github.com/graphql-go/graphql"
)

// Do use the url queries and the MainSchema in order to
// Get values from db
func Do(request *http.Request) *graphql.Result {
	var token string
	jwtcookie, cookieError := request.Cookie("jwt-token")
	if cookieError == nil {
		token = jwtcookie.Value

	}

	return graphql.Do(graphql.Params{
		Schema:        MainSchema,
		RequestString: request.URL.Query().Get("query"),
		Context:       context.WithValue(context.Background(), "token", token),
	})
}
