package error

import (
	"encoding/json"
	"errors"
	"net/http"
)

// CreateErrorFromString create and returns an error which
// have for message the given message
func CreateErrorFromString(message string) error {
	return errors.New(message)
}

func ShowOnBrowser(resp http.ResponseWriter, err error) {
	bitErr := []byte(err.Error())

	resp.WriteHeader(http.StatusBadRequest)
	resp.Write(bitErr)
}

func SendJSON(resp http.ResponseWriter, vErr []string) error {
	bs, mErr := json.Marshal(vErr)
	if mErr != nil {
		return mErr
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.Write(bs)
	return nil
}
