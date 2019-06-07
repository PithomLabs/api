package randDict

import (
	"fmt"
	"log"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

const (
	// Max password length (the password can actually be more than this number)
	maxPassLen int = 20
	// The raw github link for the dictionnary file
	githubDictionary = "https://raw.githubusercontent.com/komfy/dict/master/wordsGist"

)

// Create a slice with the words from the github link
func createWordsSlice() ([]string, int, error) {
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

// The /rand_dict handler functi
func Handler(writer http.ResponseWriter, request *http.Request) {
	// Change the rand seed every time we asked for a password
	rand.Seed(time.Now().UnixNano())
	words, wordsLen, err := createWordsSlice()
	if err != nil {
		log.Fatal(err)
		
	}

	var pass []string
	for length := 0; length < maxPassLen; {
		word := words[rand.Intn(wordsLen)]
		length += len(word)
		pass = append(pass, word)
	}
	password := strings.Join(pass, "-")
	fmt.Fprintln(writer, password)
}
