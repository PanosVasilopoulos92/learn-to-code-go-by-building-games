package main

/*
Problem 12 - Fibonacci Sequence
Write a program which takes as input a number N and generates the fibonacci sequence up to N.

In the fibonacci sequence, the nth number is equal to the sum of the previous two numbers.

The first and second numbers are always equal to 1.

*/

import (
	"fmt"
	"log"
)

func main() {
	userInput := 0
	fmt.Println("Insert an integer number:")
	_, err := fmt.Scanf("%d\n", &userInput)
	if err != nil {
		log.Fatalln(err)
	}

	result := fibonacciNumbers(userInput)
	fmt.Println(result)

}

func fibonacciNumbers(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	return fibonacciNumbers(n-1) + fibonacciNumbers(n-2)
}
