package lambdas

import (
	"net/http"
	"strconv"

	"github.com/komfy/api/internal/database"
	"github.com/komfy/api/internal/jwt"
)

func DeleteUserHandler(res http.ResponseWriter, req *http.Request) {
	token, exists := req.Header["Authorization"]
	if !exists {
		http.Error(res, "there is no authorization header", http.StatusBadRequest)
	}
	if token[0] == "" {
		http.Error(res, "authorization header is empty", http.StatusBadRequest)
	}

	userInt, err := jwt.IsValid(token[0])
	userID, err := strconv.Atoi(userInt.(map[string]string)["ID"])
	if err != nil {
		http.Error(res, "couldn't convert an id to int", http.StatusBadGateway)
	}
	user, err := database.GetUserByID(uint(userID))
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
	}
	err = database.DeleteUser(user)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
	}

}
