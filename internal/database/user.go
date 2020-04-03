package database

import (
	"github.com/komfy/api/internal/structs"
)

func IsValidUser(user *structs.User) (bool, error) {
	nUsers := 0
	cuErr := openDatabase.Instance.Table("users").Where("username = ? AND email = ?", user.Username, user.Email).Count(&nUsers).Error
	if cuErr != nil {
		return false, nil
	}

	return nUsers == 0, nil
}

func AddUser(user *structs.User) error {
	// INSERT INTO users(...) VALUES(`user`)
	auErr := openDatabase.Instance.Create(user).Error
	if auErr != nil {
		return auErr
	}
	// Add the User's ID value to its settings
	user.Settings.UserID = user.ID
	// INSERT INTO settings(...) VALUES(`user.Settings`)
	asErr := openDatabase.Instance.Create(&user.Settings).Error
	if asErr != nil {
		return asErr
	}

	return nil
}

func DeleteUser(user *structs.User) error {
	// DELETE FROM users WHERE id = `user.id`
	// User's settings, posts, and comments will be deleted
	// Thanks to ON DELETE CASCADE
	dErr := openDatabase.Instance.Delete(user).Error
	if dErr != nil {
		return dErr
	}
	return nil
}

func GetUserByName(username string) (*structs.User, error) {
	user := &structs.User{}
	// SELECT * FROM users WHERE username = `username`
	guErr := openDatabase.Instance.Where("username = ?", username).First(user).Error
	if guErr != nil {
		return nil, guErr
	}

	// SELECT * FROM settings WHERE user_id = `user.ID`
	gsErr := GetUserSettings(user)
	if gsErr != nil {
		return nil, gsErr
	}

	return user, nil
}

func GetUserByID(id uint) (*structs.User, error) {
	user := &structs.User{}
	// SELECT * FROM users WHERE id = `id`
	guErr := openDatabase.Instance.Where("user_id = ?", id).First(user).Error
	if guErr != nil {
		return nil, guErr
	}

	gsErr := GetUserSettings(user)
	if gsErr != nil {
		return nil, gsErr
	}

	return user, nil
}

func GetUserSettings(user *structs.User) error {
	// SELECT * FROM settings WHERE user_id = `user.ID`
	csErr := openDatabase.Instance.Model(user).Related(&user.Settings).Error
	if csErr != nil {
		return csErr
	}

	return nil
}

func UpdateCheck(user *structs.User) error {
	uErr := openDatabase.Instance.Model(user).UpdateColumn("checked", true).Error
	if uErr != nil {
		return uErr
	}

	return nil
}
