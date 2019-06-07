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
	// Read all the words inside the file "wordsGist"
	file, err := ioutil.ReadFile("template/wordsGist")
	// Make a get request to the github dictonary's link
	resp, err := http.Get(githubDictionary)
	if err != nil {
		log.Fatal(err)
		return nil, 0, err
	}

	// Read the github dictionary page body
	body, bErr := ioutil.ReadAll(resp.Body)
	if bErr != nil {
		log.Fatal(bErr)
		return nil, 0, bErr
	}

	// Defer the Body.Close() func and get any of the errors
	// It could return, for more info read: 
	// https://www.joeshaw.org/dont-defer-close-on-writable-files/
	defer func() {
		if bodyErr := resp.Body.Close(); err != nil {
			log.Fatal(bodyErr)
			err = bodyErr
		}
	}()

	// Split the github dictionary's body into a slice
	words := strings.Split(string(body), "\n")
	return words, len(words), err
}

// The /rand_dict handler functi
func Handler(writer http.ResponseWriter, request *http.Request) {
	// Change the rand seed every time we asked for a password
	rand.Seed(time.Now().UnixNano())
	words, wordsLen, err := createWordsSlice()
	if err != nil {
		fmt.Fprintf(writer, "Error: %v", err)
	}

	var pass []string
	for length := 0; length < maxPassLen; {
		word := words[rand.Intn(wordsLen)]
		length += len(word)
		pass = append(pass, word)
	}
	fmt.Fprintln(writer, strings.Join(pass, "-"))
}
