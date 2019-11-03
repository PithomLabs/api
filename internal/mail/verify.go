package mail

import (
	"github.com/komfy/api/internal/database"
	err "github.com/komfy/api/internal/error"
)

func Verify(code string) error {
	user := database.UserByID(code)

	if user.Username == "" {
		return err.ErrUserDoesntExist
	}

	if user.Checked {
		return err.ErrUserAlreadyChecked
	}

	database.UpdateCheck(user)

	return nil
}
