package lambdas

import (
	"net/http"
	"strconv"

	"github.com/komfy/api/internal/database"
	"github.com/komfy/api/internal/jwt"
	"github.com/komfy/api/internal/sign/register"
)

func DeleteUserHandler(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		http.Error(res, "impossible to parse form", http.StatusBadRequest)
		return
	}
	password := req.FormValue("password")

	if password == "" {
		http.Error(res, "password field is not present", http.StatusBadRequest)
		return
	}

	hashedPass, err := register.HashPassword(password)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
	}

	token, exists := req.Header["Authorization"]
	if !exists {
		http.Error(res, "there is no authorization header", http.StatusBadRequest)
		return
	}
	if token[0] == "" {
		http.Error(res, "authorization header is empty", http.StatusBadRequest)
		return
	}

	userInt, err := jwt.IsValid(token[0])
	userID, err := strconv.Atoi(userInt.(map[string]string)["ID"])
	if err != nil {
		http.Error(res, "couldn't convert an id to int", http.StatusBadGateway)
		return
	}
	user, err := database.GetUserByID(uint(userID))
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	if user.Password != hashedPass {
		http.Error(res, "passwords are not matching", http.StatusBadRequest)
		return
	}

	err = database.DeleteUser(user)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

}
