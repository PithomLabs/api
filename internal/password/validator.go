package password

import (
	"strconv"
	"unicode"

	errPack "github.com/komfy/api/internal/error"
)

type Criteria struct {
	Length   bool
	Number   bool
	Upper    bool
	Special  bool
	Position int
}

var perfect Criteria = Criteria{true, true, true, true, 0}

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

//ThrowErrors is a helper function which checks Criteria and throws a slice of errors.
func ThrowErrors(c Criteria) (errors []string) {
	if !c.Length {
		errors = append(errors, errPack.ErrShortPassword.Error())
	}
	if !c.Number {
		errors = append(errors, errPack.ErrNoNumber.Error())
	}
	if !c.Upper {
		errors = append(errors, errPack.ErrNoNumber.Error())
	}
	if !c.Special {
		errors = append(errors, errPack.ErrNoSpecial.Error())
	}
	if c.Position != 0 {
		errors = append(errors, "there is prohibited symbol on position"+strconv.Itoa(c.Position))
	}
	return errors
}
