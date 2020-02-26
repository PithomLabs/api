package error

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// CreateErrorFromString create and returns an error which
// have for message the given message
func CreateErrorFromString(message string) error {
	return errors.New(message)
}

func CreateArgumentsError(arg string, argType string) error {
	return CreateErrorFromString(
		fmt.Sprintf(
			ErrArgumentWrongTypeTemplate.Error(), arg, argType),
	)
}

func ShowOnBrowser(resp http.ResponseWriter, err error) {
	bitErr := []byte(err.Error())

	resp.WriteHeader(http.StatusBadRequest)
	resp.Write(bitErr)
}

func SendJSON(resp http.ResponseWriter, vErr []string) error {
<<<<<<< HEAD
	bs, mErr := json.Marshal(vErr)
	if mErr != nil {
		return mErr
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.Write(bs)
=======
	bs, err := json.Marshal(vErr)
	if err != nil {
		return err
	}
	resp.Header().Set("Content-Type", "application/json")
	_, err = resp.Write(bs)
>>>>>>> master
	return nil
}
