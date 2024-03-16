package main

/*
Problem 14 - Prime Checker
Write a program which takes as input a number N and generates all numbers in the range [1; N] which are prime numbers.

*/

import (
	"fmt"
	"log"
)

func main() {
	userInput := 0
	fmt.Println("Insert an integer number in order to find out if it is a primary number or not:")
	_, err := fmt.Scanf("%d\n", &userInput)
	if err != nil {
		fmt.Scanln(&userInput)
		log.Fatalln(err)
	}

	if userInput <= 1 {
		fmt.Println("Not a Prime number.")
	}

	for i := 2; i <= userInput; i++ {
		fmt.Println(i)
		isPrime := true
		// At first i will equals j so the first time it won't access the inner loop. After that the outer loop will always be one number ahead.
		for j := 2; j < i; j++ {
			if i%j == 0 {
				isPrime = false
			}
		}

		if isPrime {
			fmt.Printf("%d, ", i)
		}
	}

}
