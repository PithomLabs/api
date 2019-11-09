package graphql

import (
	"context"
	"net/http"

	"github.com/graphql-go/graphql"
)

func DoWith(schema graphql.Schema, request *http.Request) *graphql.Result {
	token, ok := request.Header["Authentication"]

	// We use that struct in order to pass multiple context variables
	cp := ContextProvider{
		Private: !ok,
		Token:   "",
	}

	if ok {
		cp.Token = token[0]
	}

	return graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: request.URL.Query().Get("query"),
		Context:       context.WithValue(context.Background(), "contextProvider", cp),
	})
}
