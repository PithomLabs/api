package mail

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/komfy/api/internal/sign"
	"github.com/komfy/api/internal/structs"
	email "gopkg.in/gomail.v2"
)

const (
	mailSubject string = "Komfy email verification"
	mailBody    string = "<h2>Komfy email verification</h2> Confirm email by clicking on this <a href='https://api.komfy.now.sh/verify?verify_code=%v'>link</a>"

	emailRegexp string = "[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,4}"
)

func Send(user *structs.User, sendChan chan sign.Transport) error {
	from := os.Getenv("user_email")
	pass := os.Getenv("pass_email")

	msg := email.NewMessage()

	msg.SetHeader("From", from)
	msg.SetHeader("To", user.Email)
	msg.SetHeader("Subject", mailSubject)
	msg.SetAddressHeader("To", user.Email, user.Username)

	dialer := email.NewDialer("smtp.gmail.com", 587, from, pass)

	<-sendChan
	msg.SetBody("text/html", fmt.Sprintf(mailBody, user.UserID))

	dErr := dialer.DialAndSend(msg)
	if dErr != nil {
		return dErr
	}

	return nil
}

func IsValid(email string) bool {
	reg, rErr := regexp.Compile(emailRegexp)
	if rErr != nil {
		log.Println(rErr)
		return false
	}

	return reg.MatchString(email)
}
