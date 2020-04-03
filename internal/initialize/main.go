package initialize

import (
	ggo "github.com/graphql-go/graphql"
	"github.com/komfy/api/internal/captcha"
	"github.com/komfy/api/internal/database"
	"github.com/komfy/api/internal/graphql"
	"github.com/komfy/api/internal/password"
)

var IsOkay bool

func TurnOkay() error {
	IsOkay = true

	captcha.InitializeMemoryStorage()

	dErr := database.InitializeDatabaseInstance()
	if dErr != nil {
		return dErr
	}

	pErr := password.GenerateWordSlice()
	if pErr != nil {
		return pErr
	}

	var sErr error
	graphql.Schema, sErr = ggo.NewSchema(ggo.SchemaConfig{
		Query: graphql.Root(),
	})
	if sErr != nil {
		return sErr
	}

	return nil
}
