package main

/*
Problem 01 - Print bigger number
Write a program which receives as input two numbers from the terminal and prints the bigger of the two.
If the numbers are equal, print “The numbers are equal”
*/

import (
	"fmt"
	"time"
)

func main() {
	num1 := 0
	num2 := 0

	fmt.Println("Please provide two numbers:")
	fmt.Scan(&num1)
	fmt.Scan(&num2)

	fmt.Printf("You provided two numbers: %d and %d\n", num1, num2)
	fmt.Println("Wait while calculation is in progress...")
	fmt.Println()
	time.Sleep(2 * time.Second)

	if num1 > num2 {
		fmt.Printf("%d is bigger than %d\n", num1, num2)
	} else if num1 == num2 {
		fmt.Printf("%d and %d are equals\n", num1, num2)
	} else {
		fmt.Printf("%d is smaller than %d\n", num1, num2)
	}
}
