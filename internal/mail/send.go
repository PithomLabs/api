package mail

import (
	"fmt"
	"os"

	"github.com/komfy/api/internal/sign"
	"github.com/komfy/api/internal/structs"
	email "gopkg.in/gomail.v2"
)

var (
	mailSubject string = "Komfy email verification"

	// HACK: yes, I know that this is shit, but it works without any Sprintf hacks
	mailBody string = "<h2>Komfy email verification</h2> Confirm email by clicking on this <a href='" + os.Getenv("frontend") + "/verify?verify_code=%v'>link</a>"
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

	// Receive an empty Transport in order to synchronize
	// with main goroutine and know when we can access
	// UserID field without trouble
	<-sendChan
	msg.SetBody("text/html", fmt.Sprintf(mailBody, user.ID))

	dErr := dialer.DialAndSend(msg)
	if dErr != nil {
		return dErr
	}

	return nil
}
