package email

import (
	"fmt"
	"log"
	"os"

	db "github.com/komfy/api/pkg/database"
	email "gopkg.in/gomail.v2"
)

// SendMail is used along side register.go in order to send
// emails for verification
func SendMail(user *db.User) {
	from := os.Getenv("user_email")
	pass := os.Getenv("pass_email")

	msg := email.NewMessage()

	msg.SetHeader("From", from)
	msg.SetHeader("To", user.Email)
	msg.SetHeader("Subject", "Komfy email verification")
	msg.SetBody("text/html", fmt.Sprintf("<h2>Komfy email verification</h2> Confirm email by clicking on this <a href='https://api.komfy.now.sh/verify?verify_code=%d'>link</a>.", user.UserID))
	msg.SetAddressHeader("To", user.Email, user.Username)

	dialer := email.NewDialer("smtp.gmail.com", 587, from, pass)

	if err := dialer.DialAndSend(msg); err != nil {
		log.Print(err)
	}
}
