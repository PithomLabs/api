package password

import (
	"math/rand"
	"strings"
	"time"
)

const maximalLength int = 20

func WordsSequence() string {
	rand.Seed(time.Now().Unix())

	var pass []string
	dictLength := len(wordSlice)

	for length := 0; length < maximalLength; {
		word := wordSlice[rand.Intn(dictLength)]
		length += len(word)
		pass = append(pass, word)
	}

	password := strings.Join(pass, "-")
	return password
}
