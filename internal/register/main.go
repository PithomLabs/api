package register

import (
	"log"
	"net/http"

	"github.com/komfy/api/internal/database"
)

// NewUser will verify user informations from the request,
// create a new database user, based on those,
// and send a mail to the user's email
func NewUser(request *http.Request) error {
	if request.Method != http.MethodPost {
		return nil // RETURN AN ERROR USING INTERNAL/ERROR
	}

	userChan := make(chan transport)
	go extractUser(request, userChan)

	dErr := doubleCheck(request)
	if dErr != nil {
		return dErr
	}

	infos := <-userChan
	close(userChan)

	if infos.Error != nil {
		return infos.Error
	}

	db, dbErr := database.Open()
	defer func() {
		cErr := db.Close()
		log.Println(cErr)
	}()

	if dbErr != nil {
		return dbErr
	}

	validChan := make(chan transport)
	tempPass := infos.User.Password
	infos.User.Password = ""

	go isValidUser(db, infos.User, validChan)

	hashed, hErr := hashPassword(tempPass)
	tempPass = ""
	if hErr != nil {
		return hErr
	}

	infos.User.Password = hashed

	userValid := <-validChan
	close(validChan)

	if !userValid.Bool {
		return userValid.Error
	}

	return nil
}
