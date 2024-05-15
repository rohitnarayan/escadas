package main

import (
	"math/rand"
	"time"
)

type Dice interface {
	Roll() int
}

type SixSideDice struct{}

func (d *SixSideDice) Roll() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(6) + 1
}
