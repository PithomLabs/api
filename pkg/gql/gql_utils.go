package gql

import (
	"github.com/graphql-go/graphql"
	err "github.com/komfy/api/pkg/error"
	"github.com/komfy/api/pkg/jwt"
)

// The type of function needed inside the privatefield function parameters
type privateFunc func(context ContextProvider) (interface{}, error)
type resolveFunc func(context ContextProvider, tokenInfos interface{}) (interface{}, error)

func generalResolveFunc(parameters graphql.ResolveParams, fn resolveFunc) (interface{}, error) {
	context, cErr := getContext(parameters)
	if cErr != nil {
		return nil, cErr
	}

	if !context.Private {
		tokenInfos, jErr := jwt.IsTokenValid(context.Token)
		if jErr != nil {
			return nil, jErr
		}

		return fn(context, tokenInfos)

	}

	return fn(context, nil)
}

// Allow us to privatise a gql field
func privatiseField(parameters graphql.ResolveParams, fn privateFunc) (interface{}, error) {
	// Get the context in order to use it
	context, cErr := getContext(parameters)
	if cErr != nil {
		return nil, cErr
	}
	// If the context is private (which mean no token was given)
	// then we return nothing, because the current field is private
	if context.Private {
		return nil, nil
	}
	// Else we return what the fn function have to return normally
	return fn(context)
}

// Return the context cast as a ContextProvider
func getContext(parameters graphql.ResolveParams) (ContextProvider, error) {
	context, ok := parameters.Context.Value("contextProvider").(ContextProvider)

	if !ok {
		return context, err.ErrContextProvider
	}
	return context, nil
}
