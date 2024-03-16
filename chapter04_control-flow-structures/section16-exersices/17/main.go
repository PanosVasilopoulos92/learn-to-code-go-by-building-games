package main

/*
Problem 17* - Guess the number Interactive
Extend the solution to Problem 07* - Guess the number, to interactively ask the user repeatedly
to guess the number until they get it right.

In the end, print the number of wrong guesses the user had.
If the number of wrong guesses is less than N/4, print “You were quite lucky!”
If the number of wrong guesses is less than N/2, print “You did alright.”
If the number of wrong guesses is more than or equal to N/2, print “It took you a while…”

*/

import (
	"fmt"
	"log"
	"math/rand"
)

func main() {
	var lowestPoint int
	var highestPoint int
	var guess int
	counter := 0

	fmt.Println("Insert two integer numbers that will correspond to the range that pseudo-random number can be generated from:")
	_, err := fmt.Scanf("%d %d\n", &lowestPoint, &highestPoint)
	if err != nil {
		fmt.Scanln(&lowestPoint, &highestPoint)
		log.Fatalln(err)
	}
	fmt.Printf("The range will be [%d;%d)\n", lowestPoint, highestPoint)
	fmt.Println("Now provide your guess number:")
	_, err = fmt.Scanf("%d\n", &guess)
	if err != nil {
		fmt.Scanln(&guess)
		log.Fatalln(err)
	}

	generatedNumber := lowestPoint + rand.Intn(highestPoint-lowestPoint)
	fmt.Println(generatedNumber)

	for {
		counter++

		if guess == generatedNumber {
			fmt.Println("You guessed the number correctly!")
			if counter < (highestPoint-lowestPoint)/4 {
				fmt.Printf("You had %d guesses. You were quite lucky.", counter)
			} else if counter < (highestPoint-lowestPoint)/2 {
				fmt.Printf("You had %d guesses. You did alright.", counter)
			} else if counter >= (highestPoint-lowestPoint)/2 {
				fmt.Printf("You had %d guesses. It took you a while…", counter)
			}
			break
		}

		fmt.Println("You guessed wrong. Try again:")
		_, err := fmt.Scanf("%d\n", &guess)
		if err != nil {
			fmt.Scanln(&guess)
			fmt.Println("Error while reading input. Try again:")
			continue
		}

	}

}
