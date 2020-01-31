package register

import (
	"net/http"

	"github.com/komfy/api/internal/database"
	err "github.com/komfy/api/internal/error"
	"github.com/komfy/api/internal/sign"
	"net/http"
	"time"
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

	/*dErr := doubleCheck(request)
	if dErr != nil {
		return dErr, nil
	}*/

	infos := <-userChan
	close(userChan)

	// Add values to Non-Null fields to user structs
	infos.User.Fullname = infos.User.Username
	infos.User.CreatedAt = uint64(time.Now().Unix())

	if infos.Error != nil {
		return infos.Error
	}

	validChan := make(chan sign.Transport)
	tempPass := infos.User.Password

	criteria := password.Validate(tempPass)
	infos.Validation = password.ThrowErrors(criteria)

	if len(infos.Validation) > 0 {
		return err.ErrPasswordNotValid, infos.Validation
	}

	infos.User.Password = ""

	go isValidUser(infos.User, validChan)

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

	database.AddUser(infos.User)
	// Send an empty Transport because the sendMail function
	// Shouldn't access the UserID field before user is added to
	// the database
	sendChan <- sign.Transport{}

	sendInfos := <-sendChan
	close(sendChan)
	if sendInfos.Error != nil {
		database.DeleteUser(infos.User)
		return sendInfos.Error
	}

	return nil
}
