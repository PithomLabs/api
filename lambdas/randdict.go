package lambdas

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	rd "github.com/komfy/api/pkg/dictutils"
)

// Max password length (the password can actually be more than this number)
const maxPassLen int = 20

// RandDictHandler corresponds to the "/rand_dict" endpoints
func RandDictHandler(writer http.ResponseWriter, request *http.Request) {
	// Change the rand seed every time we asked for a password
	rand.Seed(time.Now().UnixNano())
	words, wordsLen, err := rd.CreateWordsSlice()
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
