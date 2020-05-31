package mail

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/komfy/api/internal/database"
	err "github.com/komfy/api/internal/error"
)

func Verify(code graphql.ID) error {
	user, idErr := database.GetUserByID(string(code))
	if idErr != nil {
		return err.ErrInDatabaseOccured
	}

	if user.Username == "" {
		return err.ErrUserDoesntExist
	}

	if user.Checked {
		return err.ErrUserAlreadyChecked
	}

	uErr := database.UpdateCheck(user)
	if uErr != nil {
		return err.ErrInDatabaseOccured
	}

	return nil
}
