package main

/*
Problem 06 - Count letters
Write a program which takes in some text and counts the number of letters in it.
The counting should be case-insensitive.

*/

import (
	"fmt"
	"strings"
)

func main() {
	poem := "A whisper of wind through the trees, sunlight dappling fallen leaves, a vibrant bird takes to the air, and for a moment, all is fair..."
	m := make(map[byte]int)
	fmt.Println(poem)

	poem = strings.ToLower(poem)
	for i := 0; i < len(poem); i++ {
		if _, ok := m[poem[i]]; !ok {
			m[poem[i]]++ // will only add if key doesn't exist
		} else {
			m[poem[i]]++ // just raises the counter
		}
	}

	for k, v := range m {
		fmt.Printf("%s: %d\n", string(k), v)
	}

}
