package main

import (
	"math/rand"
	"time"
)

func aleatorionumero(maximo int) int {

	r1 := rand.New(rand.NewSource(time.Now().UnixNano()))

	return r1.Intn(maximo)
}
