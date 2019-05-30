package handler

import (
	// Go's packages
	"fmt"
	"net/http"
	"encoding/json"

	// Community's packages 
	"github.com/graphql-go/graphql"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// Creating a Schema
	graphqlFields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func (params graphql.ResolveParams) (interface{}, error) {
					return "world", nil
			},
		},
	}

	newQuery := graphql.ObjectConfig{
		Name: "MainQuery",
		Fields: graphqlFields,
	}

	schema, _err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(newQuery),
	})
	if _err != nil {
		fmt.Fprintf(w, "Unable to create the default schema: %v", _err)
	}

	query := `{
			hello
	}`

	params := graphql.Params{
		Schema: schema,
		RequestString: query,
	}

	response := graphql.Do(params)
	responseJSON, _err := json.Marshal(response)
	if _err != nil {
		fmt.Fprintf(w, "Unable to marshal the graphql response")
	}

	fmt.Fprintf(w, "%s", responseJSON)

}
