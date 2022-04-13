package main

import (
	"math/rand"
	"time"
)

func colorsText() string {

	rand.Seed(time.Now().Unix())
	colorsText := []string{"\033[31m", "\033[32m", "\033[33m", "\033[34m", "\033[35m", "\033[36m", "\033[37m"}
	n := rand.Int() % len(colorsText)
	return colorsText[n]
}
