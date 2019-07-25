package email

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	db "github.com/komfy/api/pkg/database"
)

const (
	mailBody     = "This is a verification email, here is your verification link:\n%s"
	mailTemplate = "api.komfy.now.sh/verify?verify_code=%v"
	subject      = "Email Verification"
)

// SendMail is used along side register.go in order to send
// emails for verification
func SendMail(user *db.User) {
	from := os.Getenv("user_email")

	msg := "From: " + from +
		"\nTo: " + user.Email +
		"\nSubject: " + subject +
		"\n\n" + fmt.Sprintf(mailBody, fmt.Sprintf(mailTemplate, user.UserID))

	mailError := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth(
			"",
			from,
			os.Getenv("pass_email"),
			"smtp.gmail.com",
		),
		from,
		[]string{user.Email},
		[]byte(msg))

	if mailError != nil {
		log.Fatal(mailError)

	}
}
