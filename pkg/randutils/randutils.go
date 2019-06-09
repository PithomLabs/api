package randutils

import (
	"math/rand"
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

// GeneratePassword: It's the function that is used inside 
// The Handler function from rand.go
func GeneratePassword() string {
	// Change the seed for random purpose
	rand.Seed(time.Now().UnixNano())
	// Create a byte slice of passLen size
	byteArr := make([]byte, passLen)
	for i := 0; i < passLen; {
		// rand.Int63() return a 63-bytes long integer
		// But we only want to keep the seqxByte of that number
		// And store it inside the idAnd variable
		if idAnd := int(rand.Int63() & seqxByte); idAnd < seqLen {
			// The idAnd integer can't be greater than the length
			// Of the sequence variable here
			byteArr[i] = sequence[idAnd]
			i++
		}
	}

	return string(byteArr)
}
