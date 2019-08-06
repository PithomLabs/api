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
func HandleErrorInHTTP(resp http.ResponseWriter, err interface{}) {
	resp.WriteHeader(http.StatusBadRequest)

	if message, ok := err.(string); ok {
		resp.Write([]byte(message))

	} else {
		resp.Write([]byte(err.(error).Error()))
	}
}
