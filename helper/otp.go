package helper

import "math/rand"

func MakeOtp(n int) string {
	var letter = []rune("1234567890")

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}
