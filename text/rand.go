package text

import (
	"math/rand"
	"strings"
	"time"
)

func Rand(seed []byte, length int) string {
	r := len(seed)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < length; i++ {
		sb.WriteByte(seed[rand.Intn(r)])
	}
	return sb.String()
}

// GenRandLetter generate rand letters
func RandLetters(length int) string {
	letters := []byte("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPLKJHGFDSAZXCVBNM0123456789")
	return Rand(letters, length)
}