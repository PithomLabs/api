package rand

//Package use for the endpoint /rand and /rand-dict
//Which print a random generated password on the html page
//Which could be used by /auth when creating an account

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

const (
	// All the possible characters within a password
	sequence = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789%$*!._#&"
	// The desire length
	passLen = 12
	// The sequence length
	seqLen = len(sequence)
	// The numbers of byte required for seqLen
	seqByte = 7
	// The maximum of bytes possible for a character index
	seqxByte = 1<<seqByte - 1
)

// The /rand endpoint's function
func Handler(writer http.ResponseWriter, request *http.Request) {
	rand.Seed(time.Now().UnixNano())
	byteArr := make([]byte, passLen)
	for i := 0; i < passLen; {
		if idAnd := int(rand.Int63() & seqxByte); idAnd < seqLen {
			byteArr[i] = sequence[idAnd]
			i++
		}
	}
	fmt.Fprintln(writer, string(byteArr))
}
