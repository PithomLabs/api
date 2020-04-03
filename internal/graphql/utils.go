package graphql

import (
	"github.com/graphql-go/graphql"
	err "github.com/komfy/api/internal/error"
	"github.com/komfy/api/internal/jwt"
)

// The type of function needed inside the privatefield function parameters
type returnFunction func(params graphql.ResolveParams, token map[string]string, args ...string) (interface{}, error)

/*
	resolvePublicField will parse params.Context, extract the token from it (if there is one to)
	and will call fn using fnArgs as its arguments.

	It will break if token is invalid.
*/
func resolvePublicField(params graphql.ResolveParams, fn returnFunction, fnArgs ...string) (interface{}, error) {
	context, token, cErr, tErr := getContextAndToken(params)
	if cErr != nil {
		return nil, cErr
	}

	if tErr != nil {
		// If the token was empty then it was
		// an empty requets but we shouldn't block it
		// because we are in the resolvePublicField function
		if context.Token == "" {
			return fn(params, token, fnArgs...)
		}
		// Else, it means the token wasn't valid and
		// we should block the request
		return nil, tErr
	}

	return fn(params, token, fnArgs...)
}

/*
	resolvePrivateField will do the same as resolvePublicField
	except it will break if token is empty, invalid or if context.HideInfos is set to true

	If it succeeds, fn is called using fnArgs as its arguments
*/
func resolvePrivateField(params graphql.ResolveParams, fn returnFunction, fnArgs ...string) (interface{}, error) {
	context, token, cErr, tErr := getContextAndToken(params)
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

	return fn(params, token, fnArgs...)
}

func getContextAndToken(params graphql.ResolveParams) (contextProvider, map[string]string, error, error) {
	context, ok := params.Context.Value("ContextProvider").(contextProvider)
	if !ok {
		return contextProvider{}, nil, err.ErrContextProvider, nil
	}

	token, jErr := jwt.IsValid(context.Token)
	if jErr != nil {
		return context, nil, nil, jErr
	}

	return context, token.(map[string]string), nil, nil
}
