package gql

import (
	"github.com/graphql-go/graphql"
)

// RootSchema is the schema using the root query
var RootSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: rootQuery,
})
