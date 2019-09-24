package email

import (
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
	msg.SetBody("text/html", "<h1>Komfy email verification</h1> Confirm email by clicking on this <a href='api.komfy.now.sh/verify?verify_code=%s'>link</a>.", user.userID)
	msg.SetAdressHeader("To", user.Email, user.Username)

	dialer := email.NewDialer("smtp.gmail.com", 587, from, pass)

	if err := dialer.DialAndSend(msg); err != nil {
		log.Print(err)
	}
}
