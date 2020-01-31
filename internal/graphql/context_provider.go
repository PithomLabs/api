package graphql

// ContextProvider is used along side gql in order
// to provide context to the queries
type ContextProvider struct {
	HideInfos bool
	Token     string
}
