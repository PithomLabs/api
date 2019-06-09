package dictutils

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const githubDictionary = "https://raw.githubusercontent.com/komfy/dict/master/wordsGist"

// Create a slice with the words from the github link
func CreateWordsSlice() ([]string, int, error) {
	// Make a get request to the github dictonary's link
	resp, err := http.Get(githubDictionary)
	if err != nil {
		log.Fatal(err)
		return nil, 0, err
	}

	// Read the github dictionary page body
	body, bodyErr := ioutil.ReadAll(resp.Body)
	if bodyErr != nil {
		log.Fatal(bodyErr)
		return nil, 0, bodyErr
	}

	// Defer the Body.Close() func and get any of the errors
	// It could return, for more info read:
	// https://www.joeshaw.org/dont-defer-close-on-writable-files/
	defer func() {
		if closedBodyErr := resp.Body.Close(); err != nil {
			log.Fatal(closedBodyErr)
			err = closedBodyErr
		}
	}()

	// Split the github dictionary's body into a slice
	words := strings.Split(string(body), "\r\n")
	return words, len(words), err
}
