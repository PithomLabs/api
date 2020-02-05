package database

import "github.com/komfy/api/internal/structs"

func IsValidUser(user *structs.User) bool {
	users := findSameUsers(user.Username, user.Email)

	return len(users) == 0
}

func findSameUsers(username, email string) (users []structs.User) {
	openDatabase.Instance.Where("email = ? OR username = ?", email, username).Find(&users)
	return users
}

func AddUser(user *structs.User) {
	openDatabase.Instance.Create(user)
}

func DeleteUser(user *structs.User) {
	openDatabase.Instance.Delete(user, "username = ?", user.Username)
}

func UserByName(username string) *structs.User {
	user := &structs.User{}
	openDatabase.Instance.First(&user, "username = ?", username)

	return user
}

func UserByID(userID string) *structs.User {
	user := &structs.User{}
	openDatabase.Instance.First(&user, "user_id = ?", userID)

	return user
}

func UpdateCheck(user *structs.User) {
	openDatabase.Instance.Model(&user).Update("checked", true)
}
