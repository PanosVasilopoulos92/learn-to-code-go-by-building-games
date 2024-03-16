package main

/*
Problem 07* - Guess the number
Write a program which generates a random number in a given range and the user has to guess what number your program has in mind (or in memory, to be more specific…)

It should read two input numbers from the terminal:
N - the range of numbers to use
X - the number the user has guessed

Then, the program should generate a random number in the range [0, N) and check if the number the user guessed is equal to the number the program generated.

*/

import (
	"fmt"
	"math/rand"
)

func main() {
	var lowestPoint int
	var highestPoint int
	var guess int
	// // Reads any data that may accidentally remained in the input buffer.
	// if dummyData, _ := fmt.Scanln(); dummyData >= 0 {
	// 	fmt.Println("dummy data was found. Program is now continued.")
	// }

	fmt.Println("Insert two integer numbers that will correspond to the range that pseudo-random number can be generated from:")
	fmt.Scanf("%d %d\n", &lowestPoint, &highestPoint)
	fmt.Printf("The range will be [%d;%d)\n", lowestPoint, highestPoint)
	fmt.Println("Now provide your guess number:")
	fmt.Scanf("%d\n", guess)

	generatedNumber := lowestPoint + rand.Intn(highestPoint-lowestPoint)
	fmt.Println(generatedNumber)
	// fmt.Printf("%T", generatedNumber)

	for {
		if guess == generatedNumber {
			fmt.Println("You guessed the number correctly!")
			break
		}

		fmt.Println("You guessed wrong. Try again:")
		n, err := fmt.Scanf("%d\n", &guess)
		if err != nil {
			fmt.Println("Error while reading input. Try again:")
			continue
		}

		if n != 1 {
			fmt.Println("Incorrect number of values provided. Please try again:")
			continue
		}
	}

}
