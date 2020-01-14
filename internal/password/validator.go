package password

import (
	"unicode"
)

type Criteria struct {
	Length   bool
	Number   bool
	Upper    bool
	Special  bool
	Position int
}

var specialChars []byte = []byte("%$*!._#&-'")

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
	c.Length = len(pass) >= passLen
	for key, char := range pass {
		switch {
		case unicode.IsNumber(char):
			c.Number = true
		case unicode.IsUpper(char):
			c.Upper = true
		case unicode.IsSymbol(char) || unicode.IsPunct(char):
			if !checkChars(char) {
				c.Position = key + 1
				c.Special = false
				return
			}
			c.Special = true
		case unicode.IsLetter(char):
			//just to check letters
		}
	}

	return
}
