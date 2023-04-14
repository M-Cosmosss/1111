package main

import (
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	i := rand.Intn(2)
	if i == 0 {
		os.Exit(-1)
	}
	return
}
