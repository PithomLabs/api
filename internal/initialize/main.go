package initialize

import (
	"github.com/komfy/api/internal/captcha"
	"github.com/komfy/api/internal/database"
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

	return nil
}
