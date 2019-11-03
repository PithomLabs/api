package password

import (
	"io/ioutil"
	"net/http"
	"strings"
)

const githubDictionary = "https://raw.githubusercontent.com/komfy/dict/master/wordsGist"

var wordSlice []string

func GenerateWordSlice() (mErr error) {
	resp, mErr := http.Get(githubDictionary)
	if mErr != nil {
		return mErr
	}

	body, mErr := ioutil.ReadAll(resp.Body)
	if mErr != nil {
		return mErr
	}

	// Defer the Body.Close() func and get any of the errors
	// It could return, for more info read:
	// https://www.joeshaw.org/dont-defer-close-on-writable-files/
	defer func() {
		if cErr := resp.Body.Close(); mErr == nil {
			mErr = cErr
		}
	}()

	wordSlice = strings.Split(string(body), "\r\n")

	return
}
