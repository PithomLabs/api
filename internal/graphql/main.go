package graphql

import (
	"context"
	"net/http"

	"github.com/graphql-go/graphql"
)

var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: Root(),
})

func Do(request *http.Request) *graphql.Result {
	token, tokenExists := request.Header["Authentication"]

	// We use that struct in order to pass multiple context variables
	cp := ContextProvider{
		HideInfos: !tokenExists,
		Token:     "",
	}

	if tokenExists {
		cp.Token = token[0]
	}

	return graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: request.URL.Query().Get("query"),
		Context:       context.WithValue(context.Background(), "ContextProvider", cp),
	})
}
