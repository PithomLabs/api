package gql

import (
	"context"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/komfy/api/pkg/database"
)

// Do use the url queries and the MainSchema in order to
// Get values from db
func Do(request *http.Request) *graphql.Result {
	var token string
	jwtcookie, cookieError := request.Cookie("jwt-token")
	if cookieError == nil {
		token = jwtcookie.Value

	} else {
		return nil
	}

	// We use that struct in order to pass multiple context variables
	cp := ContextProvider{
		Token:    token,
		Database: database.OpenDatabase(),
	}

	// Close the database after gql has done all of its ResolveFunc
	defer cp.Database.CloseDB()

	return graphql.Do(graphql.Params{
		Schema:        MainSchema,
		RequestString: request.URL.Query().Get("query"),
		Context:       context.WithValue(context.Background(), "context_provider", cp),
	})
}
