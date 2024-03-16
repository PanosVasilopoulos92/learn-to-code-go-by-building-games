package main

/*
Problem 09* - Number as word
Write a program which takes an input number in the range [0; 1000) and prints it as a word.

*/

// Solution in progress, code not fully functional...

import (
	"fmt"
	"log"
)

func main() {
	var inputNumber int
	var tmpNumber int
	var remainder int
	var lastDigit string
	var secondLastDigit string
	var thirdLastDigit string
	var firstDigit string
	var numberInWords string

	fmt.Println("Insert the number of your choice:")
	n, err := fmt.Scanf("%d\n", &inputNumber)
	if err != nil {
		log.Fatal(err)
	}

	if n != 1 {
		log.Fatalln("More than one arguments were provided. Can have only one.")
	}

	tmpNumber = inputNumber
	remainder = tmpNumber
	counter := 0

	for {
		tmpNumber /= 10
		counter++
		fmt.Println(tmpNumber)
		if tmpNumber <= 0 {
			break
		}
	}

	tmpNumber = inputNumber
	remainder = tmpNumber

	for remainder != 0 {
		remainder = tmpNumber % 10
		tmpNumber /= 10

		switch remainder {
		case 1:
			lastDigit = "one"
		case 2:
			lastDigit = "two"
		case 3:
			lastDigit = "three"
		case 4:
			lastDigit = "four"
		case 5:
			lastDigit = "five"
		case 6:
			lastDigit = "six"
		case 7:
			lastDigit = "seven"
		case 8:
			lastDigit = "eight"
		case 9:
			lastDigit = "nine"
		}

		remainder = tmpNumber % 10
		tmpNumber /= 10

		if remainder == 0 && counter == 1 {
			break
		}
		// fmt.Println(remainder)
		// fmt.Println(tmpNumber)

		switch remainder {
		case 2:
			secondLastDigit = "twenty"
		case 3:
			secondLastDigit = "thirty"
		case 4:
			secondLastDigit = "forty"
		case 5:
			secondLastDigit = "fifty"
		case 6:
			secondLastDigit = "sixty"
		case 7:
			secondLastDigit = "seventy"
		case 8:
			secondLastDigit = "eighty"
		case 9:
			secondLastDigit = "ninety"
		}

		remainder = tmpNumber % 10
		tmpNumber /= 10

		if remainder == 0 && counter == 2 {
			break
		}
		// fmt.Println(remainder)
		// fmt.Println(tmpNumber)

		switch remainder {
		case 1:
			thirdLastDigit = "one hundred"
		case 2:
			thirdLastDigit = "two hundred"
		case 3:
			thirdLastDigit = "three hundred"
		case 4:
			thirdLastDigit = "four hundred"
		case 5:
			thirdLastDigit = "five hundred"
		case 6:
			thirdLastDigit = "six hundred"
		case 7:
			thirdLastDigit = "seven hundred"
		case 8:
			thirdLastDigit = "eight hundred"
		case 9:
			thirdLastDigit = "nine hundred"
		}

		remainder = tmpNumber % 10
		tmpNumber /= 10

		if remainder == 0 && counter == 3 {
			break
		}

		switch remainder {
		case 1:
			firstDigit = "one thousand"
		default:
			log.Fatal(err)
		}
	}

	switch counter {
	case 1:
		numberInWords = lastDigit
		fmt.Println(numberInWords)
	case 2:
		numberInWords = secondLastDigit + " " + lastDigit
		fmt.Println(numberInWords)
	case 3:
		numberInWords = thirdLastDigit + " " + secondLastDigit + " " + lastDigit
		fmt.Println(numberInWords)
	case 4:
		numberInWords = firstDigit + " " + thirdLastDigit + " " + secondLastDigit + " " + lastDigit
		fmt.Println(numberInWords)
	}

}
