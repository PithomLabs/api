package lambdas

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
	"time"
	"math/rand"
)

const (
	maxPassLen = 20
	
)

// The /rand_dict handler function
func Handler(writer http.ResponseWriter, req *http.Request) {
	rand.Seed(time.Now().UnixNano())
	words, wordsLen, err := createWordsSplit()
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	var pass []string
	var length int
	for length < maxPassLen {
		word := words[rand.Intn(wordsLen)]
		length += len(word)
		pass = append(pass, word)
	}
	fmt.Fprintln(writer, strings.Join(pass, "-"))
}

func createWordsSplit() ([]string, int, error) {
	file, err := ioutil.ReadFile("../wordsGist")
	if err != nil {
		return nil, 0, err
	}
	words := strings.Split(string(file), "\n")
	return words, len(words), nil
}