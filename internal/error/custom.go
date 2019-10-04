package error

import "errors"

var (
	// ErrValueMissing is used inside reg.go and auth.go
	ErrValueMissing = errors.New("a value is missing from the request data")
	// ErrValueMissingTemplate is used inside reg.go and auth.go
	ErrValueMissingTemplate = "sorry but you forgot ur %s"
	// ErrUserNotValid is used in reg.go and auth.go
	ErrUserNotValid = errors.New("those credentials are already used")
	// ErrBadPassword is used in auth.go
	ErrBadPassword = errors.New("given password does not match with db")
	// ErrUserIsntCheck is used in auth.go
	ErrUserIsntCheck = errors.New("the given user has not check is email address")
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
	// ErrCaptchaHeaderMissing is used in captcha.go
	ErrCaptchaHeaderMissing = errors.New("the captcha header is missing")
	// ErrDigitsMissing is used in captcha.go
	ErrDigitsMissing = errors.New("digits are missing")
	// ErrCaptchaInvalid is used in captcha.go
	ErrCaptchaInvalid = errors.New("captcha is invalid, try again")
	// ErrContentTypeMissing is used inside authentication.go
	ErrContentTypeMissing = errors.New("content-type header is missing")
	// ErrMultipartFormData is used inside authentication.go
	ErrMultipartFormData = errors.New("the multipart/form-data doesn't have a boundary")
	// ErrContextProvider is used inside gql_resolve_funcs.go
	ErrContextProvider = errors.New("the context could not been cast as ContextProvider")
	// ErrDoubleCheck is used inside register/utils.go
	ErrDoubleCheck = errors.New("double check failed")
)
