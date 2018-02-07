package server

import (
	"fmt"
	"math/rand"
)

const (
	Alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func RandomPassword() string {
	return fmt.Sprintf("%s%s%s%s", randomChar(), randomChar(), randomChar(), randomChar())
}

func randomChar() string {
	return string(Alphabet[rand.Intn(26)])
}
