package main

import "fmt"

/*
Problem 11 - Factorial
Write a program which takes as input a number N and calculates N! (factorial).

The factorial of a number is the multiplication of the numbers [1; N]

*/

func main() {
	userInput := 0
	fmt.Println("Insert number to learn it's factorial:")
	fmt.Scanf("%d\n", &userInput)

	factorial := factorial(userInput)
	fmt.Printf("Factorial of %d is %d\n", userInput, factorial)
}

func factorial(n int) int {
	result := 1
	if n == 0 {
		return result
	}

	result = n * factorial(n-1)
	return result
}
