package password

import (
	"unicode"
)

type Criteria struct {
	length   bool
	number   bool
	upper    bool
	special  bool
	position int
}

var specialChars []byte = []byte("%$*!._#&")

func checkChars(c rune) bool {
	for _, v := range specialChars {
		if c == rune(v) {
			return true
		}
	}
	return false
}

//Validate provides server side password validation
func Validate(pass string) (c Criteria) {
	c.length = len(pass) >= passLen
	for key, char := range pass {
		switch {
		case unicode.IsNumber(char):
			c.number = true
		case unicode.IsUpper(char):
			c.upper = true
		case unicode.IsSymbol(char) || unicode.IsPunct(char):
			if !checkChars(char) {
				c.position = key + 1
				c.special = false
				return
			}
			c.special = true
		case unicode.IsLetter(char):
			//just to check letters
		}
	}

	return
}
