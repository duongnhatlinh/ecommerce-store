package helper

import (
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSequence(n int) string {
	b := make([]rune, n)

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for i := range b {
		b[i] = letters[r1.Intn(9999)%len(letters)]
	}

	return string(b)
}

func GenerateSalt(length int) string {
	if length < 0 {
		length = 25
	}
	return randSequence(length)
}
