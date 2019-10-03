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
	token, ok := request.Header["Authentication"]

	// We use that struct in order to pass multiple context variables
	cp := ContextProvider{
		Private:  !ok,
		Token:    "",
		Database: database.OpenDatabase(),
	}

	if ok {
		cp.Token = token[0]
	}

	// Close the database after gql has done all of its ResolveFunc
	defer cp.Database.CloseDB()

	return graphql.Do(graphql.Params{
		Schema:        RootSchema,
		RequestString: request.URL.Query().Get("query"),
		Context:       context.WithValue(context.Background(), "contextProvider", cp),
	})
}
