package initialize

import (
	"io/ioutil"
	"math/rand"
	"time"

	ggo "github.com/graph-gophers/graphql-go"
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

	schema, err := ioutil.ReadFile("./internal/graphql/schema.graphql")
	if err != nil {
		iErrs = append(iErrs, err)
	}
	var sErr error
	graphql.Schema, sErr = ggo.ParseSchema(string(schema), &graphql.RootResolver{})
	if sErr != nil {
		iErrs = append(iErrs, sErr)
	}

	return iErrs
}
