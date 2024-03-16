package main

/*
Problem 10 - Count to N
Write a program which takes as input a number N and then generates the numbers from [0; N).

*/

import (
	"fmt"
	"log"
)

func main() {
	var userInput int
	fmt.Println("Provide an integer number:")
	n, err := fmt.Scanf("%d\n", &userInput)
	if err != nil {
		fmt.Scanln(&userInput)
		log.Fatalln("Error while reading input.", err)
	}

	if n != 1 {
		log.Fatalln("Cannot provide more than one arguments.")
	}

	for i := 0; i < userInput; i++ {
		fmt.Println(i)
	}
}
