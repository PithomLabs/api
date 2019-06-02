package randDict

import (
	"fmt"
	"net/http"
	//"io/ioutil"
	//"strings"
	//"time"
	//"math/rand"
)

const (
	maxPassLen = 20
	
)
/*
// Create a slice with the words from the words' file
func createWordsSlice() ([]string, int, error) {
	// Read all the words inside the file "wordsGist"
	file, err := ioutil.ReadFile("../wordsGist")
	if err != nil {
		return nil, 0, err
	}
	// Split the words into a slice
	words := strings.Split(string(file), "\n")
	return words, len(words), nil
}
*/
// The /rand_dict handler function
func Handler(writer http.ResponseWriter, req *http.Request) {
	// Change the rand seed every time we asked for a password
	/*
	rand.Seed(time.Now().UnixNano())
	words, wordsLen, err := createWordsSplit()
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
	*/
	fmt.Fprintln(writer, "TEST")
}