package error

import (
	"errors"
	"net/http"
)

// CreateErrorFromString create and returns an error which
// have for message the given message
func CreateErrorFromString(message string) error {
	return errors.New(message)
}

func ShowOnBrowser(resp http.ResponseWriter, err error) {
	resp.WriteHeader(http.StatusBadRequest)

	bitErr := []byte(err.Error())
	resp.Write(bitErr)
}
