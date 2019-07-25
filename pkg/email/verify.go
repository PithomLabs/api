package email

import (
	db "github.com/komfy/api/pkg/database"
	err "github.com/komfy/api/pkg/error"
)

// VerifyUser use the code in order to verify a user
func VerifyUser(code string) error {
	user := db.AskUserByID(code)

	if user.Checked {
		return err.ErrUserAlreadyChecked

	}

	db.UpdateCheckValue(user)

	return nil
}
