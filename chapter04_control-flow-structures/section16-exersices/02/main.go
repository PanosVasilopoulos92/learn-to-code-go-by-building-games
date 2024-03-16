package main

/*
Problem 02 - Number is in range
Write a program which checks if a given input number x is in the range [a; b). That is a inclusively and b exclusively.

x, a and b are given as input in the terminal.

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
	// I used bufio package instead of the simpler fmt.Scanf() for exploring/practicing purposes.
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Please the range by inserting the lowest point first and the highest point after:")
	scanner.Scan()
	a := scanner.Text()
	aInt, err := strconv.Atoi(strings.TrimSpace(a))
	if err != nil {
		log.Fatalln("error while parsing sting to integer.")
	}

	scanner.Scan()
	b := scanner.Text()
	bInt, err := strconv.Atoi(strings.TrimSpace(b))
	if err != nil {
		log.Fatalln("error while parsing sting to integer.")
	}
	fmt.Printf("The range is %v inclusive - %v exclusive\n", a, b)

	fmt.Println("Now provide the number you want to check if in range above:")
	scanner.Scan()
	x := scanner.Text()
	xInt, err := strconv.Atoi(strings.TrimSpace(x))
	if err != nil {
		log.Fatalln("error while parsing sting to integer.")
	}

	fmt.Printf("Please wait while we check if %v is in range of [%v;%v)...\n", xInt, aInt, bInt)
	time.Sleep(2 * time.Second)
	fmt.Println()

	if xInt >= aInt && xInt < bInt {
		fmt.Printf("%v is in range of [%v;%v)", xInt, aInt, bInt)
	} else {
		fmt.Println("Number not in range.")
	}

}
