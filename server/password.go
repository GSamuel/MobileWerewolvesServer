package server

import (
	"fmt"
	"github.com/google/uuid"
	"math/rand"
)

const (
	Alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func RandomCode() string {
	return fmt.Sprintf("%s%s%s%s", randomChar(), randomChar(), randomChar(), randomChar())
}

func RandomId() string {
	return uuid.New().String()
}

func randomChar() string {
	return string(Alphabet[rand.Intn(26)])
}
