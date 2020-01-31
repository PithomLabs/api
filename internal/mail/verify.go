package mail

import (
	"github.com/komfy/api/internal/database"
	err "github.com/komfy/api/internal/error"
)

func Verify(code string) error {
	user, idErr := database.GetUserByID(code)
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
