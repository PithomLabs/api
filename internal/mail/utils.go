package mail

import (
	"log"
	"regexp"
)

const emailRegexp string = "[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,4}"

func IsValid(email string) bool {
	reg, rErr := regexp.Compile(emailRegexp)
	if rErr != nil {
		log.Println(rErr)
		return false
	}

	return reg.MatchString(email)
}
