package graphql

import (
	"github.com/graphql-go/graphql"
	err "github.com/komfy/api/internal/error"
	"github.com/komfy/api/internal/jwt"
)

// The type of function needed inside the privatefield function parameters
type returnFunction func(token interface{}) (interface{}, error)

func resolvePublicField(parameters graphql.ResolveParams, fn returnFunction) (interface{}, error) {
	context, token, cErr, tErr := getContextAndToken(parameters)
	if cErr != nil {
		return nil, cErr
	}

	if tErr != nil {
		// If the token was empty then it was
		// an empty requets but we shouldn't block it
		// because we are in the resolvePublicField function
		if context.Token == "" {
			return fn(token)
		}
		// Else, it means the token wasn't valid and
		// we should block the request
		return nil, tErr
	}

	return fn(token)
}

func resolvePrivateField(parameters graphql.ResolveParams, fn returnFunction) (interface{}, error) {
	context, token, cErr, tErr := getContextAndToken(parameters)
	if cErr != nil {
		return nil, cErr
	}

	if tErr != nil {
		return nil, tErr
	}
	// HideInfos contains a boolean telling us if, or not,
	// we should give the user those informations.
	// Here we are in a private field, so we shoudln't give informations
	// if HideInfos == true
	if context.HideInfos {
		return nil, nil
	}

	return fn(token)
}

func getContextAndToken(parameters graphql.ResolveParams) (contextProvider, interface{}, error, error) {
	context, ok := parameters.Context.Value("ContextProvider").(contextProvider)
	if !ok {
		return contextProvider{}, nil, err.ErrContextProvider, nil
	}

	token, jErr := jwt.IsValid(context.Token)
	if jErr != nil {
		return context, nil, nil, jErr
	}

	return context, token, nil, nil
}
