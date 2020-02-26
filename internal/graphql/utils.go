package graphql

import (
	"github.com/graphql-go/graphql"
	err "github.com/komfy/api/internal/error"
	"github.com/komfy/api/internal/jwt"
)

// The type of function needed inside the privatefield function parameters
type returnFunction func(context ContextProvider, token interface{}) (interface{}, error)

func generalResolveFunc(parameters graphql.ResolveParams, fn returnFunction) (interface{}, error) {
	// We get the ContextProvider and the jwt token
	context, token, ctErr := getContextAndToken(parameters)
	if ctErr != nil {
		return nil, ctErr
	}

	// Here, it doesn't matter if the token is given or not,
	// so we just pass it to the returnFunction
	return fn(context, token)
}

// Allow us to privatise a gql field
func privatiseField(parameters graphql.ResolveParams, fn returnFunction) (interface{}, error) {
	// Get the context and token in order to use them
	context, token, ctErr := getContextAndToken(parameters)
	if ctErr != nil {
		return nil, ctErr
	}
	// If the context is private (which mean no token was given)
	// then we return nothing, because the current field is private
	if context.HideInfos {
		return nil, nil
	}
	// Else we return what the returnFunction has to return normally
	return fn(context, token)
}

func getContextAndToken(parameters graphql.ResolveParams) (ContextProvider, interface{}, error) {
	context, ok := parameters.Context.Value("ContextProvider").(ContextProvider)
	if !ok {
		return ContextProvider{}, nil, err.ErrContextProvider
	}

	token, jErr := jwt.IsValid(context.Token)
	if jErr != nil {
		return ContextProvider{}, nil, jErr
	}

	return context, token, jErr
}
