package gql

import (
	"github.com/graphql-go/graphql"
)

var mainSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: rootQuery,
})
