package error

import (
	"errors"
	"net/http"
)

var (
	// ErrValueMissing is used inside reg.go and auth.go
	ErrValueMissing = errors.New("value is missing")
	// ErrValueMissingTemplate is used inside reg.go and auth.go
	ErrValueMissingTemplate = "sorry but you forgot ur %s"
	// ErrUserNotValid is used in reg.go and auth.go
	ErrUserNotValid = errors.New("those credentials are already used")
	// ErrBadPassword is used in auth.go
	ErrBadPassword = errors.New("given password does not match with db")
	// ErrHashing is used in reg.go
	ErrHashing = errors.New("An error occured while trying to hash password")
	// ErrUserAlreadyChecked is used in verify.go
	ErrUserAlreadyChecked = errors.New("user is already checked")
	// ErrTokenForgotten is used in jwt.go
	ErrTokenForgotten = errors.New("token is missing")
	// ErrSigningMethod is used in jwt.go
	ErrSigningMethod = errors.New("signing method wasn't matching")
	// ErrTokenNotValid is used in jwt.go
	ErrTokenNotValid = errors.New("token is not valid")
)

// HandleErrorInHTTP is used in order to write messages
// On api webpage when an error occurs
func HandleErrorInHTTP(resp http.ResponseWriter, err error) {
	resp.WriteHeader(http.StatusBadRequest)

	resp.Write([]byte(err.Error()))
}

// CreateError create a new error based on the given string message
func CreateError(errorMessage string) error {
	return errors.New(errorMessage)
}
