package helper

import (
	"math/rand"
	"strings"
	"time"
)

func RandomString(length int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, length)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

func UnixRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("C2DBEWK6SX8PQ9R3FG4HA1YZLM7NOI5J0TUV" + "luk9vmw6ob1cd2xenf3gp7qr8sah4i5tj0yz" + "0123456789")
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	str := b.String()
	return str
}
