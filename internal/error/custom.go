package error

import "errors"

var (
	ErrValueMissing         = errors.New("a value is missing from the request data")
	ErrValueMissingTemplate = "sorry but you forgot ur %s"
	ErrUserNotValid         = errors.New("those credentials are already used")
	ErrBadPassword          = errors.New("given password does not match with db")
	ErrUserIsntCheck        = errors.New("the given user has not check is email address")
	ErrHashing              = errors.New("An error occured while trying to hash password")
	ErrUserAlreadyChecked   = errors.New("user is already checked")
	ErrUserDoesntExist      = errors.New("this user does not exist")
	ErrPostDoesntExist      = errors.New("this post does not exist")
	ErrTokenForgotten       = errors.New("token is missing")
	ErrSigningMethod        = errors.New("signing method wasn't matching")
	ErrTokenNotValid        = errors.New("token is not valid")
	ErrCaptchaHeaderMissing = errors.New("the captcha header is missing")
	ErrDigitsMissing        = errors.New("digits are missing")
	ErrCaptchaInvalid       = errors.New("captcha is invalid, try again")
	ErrContentTypeMissing   = errors.New("content-type header is missing")
	ErrMultipartFormData    = errors.New("the multipart/form-data doesn't have a boundary")
	ErrContextProvider      = errors.New("the context could not been cast as ContextProvider")
	ErrDoubleCheck          = errors.New("double check failed")
	ErrMethodNotValid       = errors.New("you try to use a non-valid method")
	ErrMailNotValid         = errors.New("your email address is not valid")
)
