package graphql

import (
	"context"
	"net/http"

	"github.com/graphql-go/graphql"
)

// Do is a wrapper of the graphql-go's Do function
func Do(request *http.Request) *graphql.Result {
	var schema, sErr = graphql.NewSchema(graphql.SchemaConfig{
		Query: root,
	})

	if sErr != nil {
		panic(sErr)
	}

	token, tokenExists := request.Header["Authentication"]

	// We use that struct in order to pass multiple context variables
	cp := contextProvider{
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
