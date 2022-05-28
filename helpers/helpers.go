package helpers

import (
	"math/rand"
	"time"
)

type SomeType struct {
	TypeName   string
	TypeNumber int
}

func RandomNumber(n int) int {
	// это нужно просто чтобы возращало ранд. число
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(n)

	return number
}
