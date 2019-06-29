package gql

import (
	"context"
	"net/http"

	"github.com/graphql-go/graphql"
)

// Do use the url queries and the MainSchema in order to
// Get values from db
func Do(request *http.Request) *graphql.Result {
	queries := request.URL.Query()
	return graphql.Do(graphql.Params{
		Schema:        MainSchema,
		RequestString: queries.Get("query"),
		Context:       context.WithValue(context.Background(), "token", queries.Get("token")),
	})
}
