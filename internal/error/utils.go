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
	resp.WriteHeader(http.StatusBadRequest)

	bitErr := []byte(err.Error())
	resp.Write(bitErr)
}

func SendJSON(resp http.ResponseWriter, vErr []string) error {
	bs, err := json.Marshal(vErr)
	if err != nil {
		return err
	}
	resp.Header().Set("Content-Type", "application/json")
	_, err = resp.Write(bs)
	return nil
}
