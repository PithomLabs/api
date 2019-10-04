package error

import "errors"

// CreateErrorFromString create and returns an error which
// have for message the given message
func CreateErrorFromString(message string) error {
	return errors.New(message)
}
