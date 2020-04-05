package initialize

import (
	"math/rand"
	"time"

	ggo "github.com/graphql-go/graphql"
	"github.com/komfy/api/internal/captcha"
	"github.com/komfy/api/internal/database"
	"github.com/komfy/api/internal/graphql"
	"github.com/komfy/api/internal/password"
)

var IsOkay bool

func TurnOkay(isDev bool) []error {
	var iErrs []error = nil

	IsOkay = true
	// Gives to rand.Seed an unique value so rand's function will
	// generate different pseudo-random numbers
	rand.Seed(time.Now().UnixNano())

	captcha.InitializeMemoryStorage()

	dErr := database.InitializeDatabaseInstance(isDev)
	if dErr != nil {
		iErrs = append(iErrs, dErr)
	}

	pErr := password.GenerateWordSlice()
	if pErr != nil {
		iErrs = append(iErrs, pErr)
	}

	var sErr error
	graphql.Schema, sErr = ggo.NewSchema(ggo.SchemaConfig{
		Query: graphql.Root(),
	})
	if sErr != nil {
		iErrs = append(iErrs, sErr)
	}

	return iErrs
}
