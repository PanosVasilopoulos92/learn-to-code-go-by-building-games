package main

/*
Problem 03* - Number is in range level 2
Write a program which checks if a given input number x is in a given range. The bounds of the range are given as the flags
in/ex which specify if the ranges are inclusive or not.

The output should be formatted as {x} is (not) in the range {range} where range has different brackets based on the in/ex flags.

For example, if the flags given are in + ex, then the range would be [a; b).
If they are in in, the range would be [a; b].

*/

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// I used bufio package instead of the simpler fmt.Scan() for practicing purposes.
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Please specify the range by inserting the lowest point first and the highest point after:")

	// Reads the first value provided by user(lowest range point)
	scanner.Scan()
	a := scanner.Text()
	aInt, err := strconv.Atoi(strings.TrimSpace(a))
	if err != nil {
		log.Fatalln("error while parsing sting to integer.")
	}

	// Reads the second value provided by user(highest range point)
	scanner.Scan()
	b := scanner.Text()
	bInt, err := strconv.Atoi(strings.TrimSpace(b))
	if err != nil {
		log.Fatalln("error while parsing sting to integer.")
	}

	fmt.Println("Now choose if you want the range to be inclusive-inclusive or inclusive-exclusive. " +
		"For inclusive-inclusive press '1', for inclusive-exclusive press '2':")
	scanner.Scan()
	typeOfRange := scanner.Text()

	switch typeOfRange {
	case "1":
		fmt.Printf("The range is %v inclusive - %v inclusive\n", a, b)
	case "2":
		fmt.Printf("The range is %v inclusive - %v exclusive\n", a, b)
	default:
		log.Fatalln("Wrong input. Program exits...")
	}

	fmt.Println("Now provide the number you want to check if in range above:")

	// Reads the third value provided by user(the actual number)
	scanner.Scan()
	x := scanner.Text()
	xInt, err := strconv.Atoi(strings.TrimSpace(x))
	if err != nil {
		log.Fatalln("error while parsing sting to integer.")
	}

	// If not 1 will be 2 or else the program would had already exited with error.
	if typeOfRange == "1" {
		fmt.Printf("Please wait while we check if %d is in range of [%d;%d]...\n", xInt, aInt, bInt)
		time.Sleep(2 * time.Second)
		fmt.Println()

		if xInt >= aInt && xInt <= bInt {
			fmt.Printf("%d is in range of [%d;%d)", xInt, aInt, bInt)
		} else {
			fmt.Println("Number not in range.")
		}
	} else {
		fmt.Printf("Please wait while we check if %d is in range of [%d;%d)...\n", xInt, aInt, bInt)
		time.Sleep(2 * time.Second)
		fmt.Println()

		if xInt >= aInt && xInt < bInt {
			fmt.Printf("%d is in range of [%d;%d)", xInt, aInt, bInt)
		} else {
			fmt.Println("Number not in range.")
		}
	}

}
