package register

import (
	"log"
	"net/http"

	"github.com/komfy/api/internal/database"
	err "github.com/komfy/api/internal/error"
	"github.com/komfy/api/internal/sign"
)

// NewUser will verify user informations from the request,
// create a new database user, based on those,
// and send a mail to the user's email
func NewUser(request *http.Request) error {
	if request.Method != http.MethodPost {
		return err.ErrMethodNotValid
	}

	userChan := make(chan sign.Transport)
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
		if cErr != nil {
			log.Println(cErr)
		}
	}()

	if dbErr != nil {
		return dbErr
	}

	validChan := make(chan sign.Transport)
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

	sendChan := make(chan sign.Transport)
	go sendMail(infos.User, sendChan)

	db.AddUser(infos.User)
	sendChan <- sign.CreateBoolTransport(true)

	sendInfos := <-sendChan
	close(sendChan)
	if sendInfos.Error != nil {
		db.DeleteUser(infos.User)
		return sendInfos.Error
	}

	return nil
}
