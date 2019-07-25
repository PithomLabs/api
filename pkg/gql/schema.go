package gql

import (
	"github.com/graphql-go/graphql"
)

// MainSchema is the schema used at /graphql
var MainSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: rootQuery,
})
