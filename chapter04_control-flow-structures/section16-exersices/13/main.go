package main

/*
Problem 13 - Number divisible by 7 and 13
Write a program which takes as input a number N and outputs all numbers in the range [1; N] which are divisible by both 7 and 13.

*/

import (
	"fmt"
	"log"
)

func main() {
	userInput := 0
	fmt.Println("Provide an integer number:")
	_, err := fmt.Scan(&userInput)
	if err != nil {
		log.Fatalln(err)
	}

	for i := 0; i <= userInput; i++ {
		if (i%7 == 0) && (i%13 == 0) {
			fmt.Println(i)
		}
	}
}
