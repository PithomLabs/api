package password

import (
	"math/rand"
	"strconv"
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
		randNum := strconv.Itoa(rand.Intn(9))
		word = strings.Title(word)
		pass = append(pass, word+randNum)
	}

	password := strings.Join(pass, "-")
	return password
}
