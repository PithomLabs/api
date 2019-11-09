package graphql

// ContextProvider is used along side gql in order
// to provide context to the queries
type ContextProvider struct {
	Private bool
	Token   string
}
