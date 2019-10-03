package gql

import "github.com/komfy/api/pkg/database"

// ContextProvider is used along side gql in order
// to provide context to the queries
type ContextProvider struct {
	Private  bool
	Token    string
	Database *database.DB
}
