package main

/*
Problem 16* - Pair multiplication
Write a program which takes as input two numbers - N and X and find all pairs in the range [1; N],
which when multiplied, produce X. Do NOT allow duplicate values.

*/

import (
	"fmt"
	"log"
)

func main() {
	userInputN := 0
	userInputX := 0

	fmt.Println("Provide two integer numbers:")
	_, err := fmt.Scanf("%d %d\n", &userInputN, &userInputX)
	if err != nil {
		// Cleans standard input in case an error occurs.
		fmt.Scanln(&userInputN, &userInputX)
		log.Fatalln(err)
	}

	if userInputN == 0 || userInputX == 0 {
		log.Fatalln("Cannot accept zero values.")
	}

	for i := 1; i <= userInputN; i++ {
		for j := 1; j <= i; j++ {
			if i*j == userInputX {
				fmt.Printf("%d * %d == %d\n", i, j, userInputX)
			}
		}
	}
}
