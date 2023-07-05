package generator

import (
	"math/rand"
	"time"
)

const allowedChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const allowedCharsLength = len(allowedChars)

func generateRandomString(minLength, maxLength int) string {
	rand.Seed(time.Now().UnixNano())

	length := rand.Intn(maxLength-minLength+1) + minLength
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = allowedChars[rand.Intn(allowedCharsLength)]
	}

	return string(result)
}
