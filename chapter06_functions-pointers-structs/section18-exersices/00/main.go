package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Create a new random number generator
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	number := random.Intn(2) + 2
	fmt.Println(number)
}
