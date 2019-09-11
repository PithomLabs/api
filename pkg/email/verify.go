package email

import (
	db "github.com/komfy/api/pkg/database"
	err "github.com/komfy/api/pkg/error"
)

// VerifyUser use the code in order to verify a user
func VerifyUser(code string) error {
	database := db.OpenDatabase()
	defer database.CloseDB()

	user := database.AskUserByID(code)

	if user.Checked {
		return err.ErrUserAlreadyChecked

	}

	database.UpdateCheckValue(user)

	return nil
}
