package random

import (
	"math/rand"
	"time"
)

func NewRandomString(size int) string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, size)
	for i := range b {
		b[i] = chars[rnd.Intn(len(chars))]
	}

	return string(b)
}
