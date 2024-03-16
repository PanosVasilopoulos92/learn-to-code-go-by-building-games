package main

/*
Problem 01 - Raise all by 2
Write a program which reads a set of integers, on a single line, stores the numbers in a slice & raises all elements’ power by 2.

*/

import (
	"fmt"
	"log"
)

func main() {
	var a, b, c, d, e int
	s := []int{}
	fmt.Println("Provide 5 integer numbers of your choice, separated by space:")
	_, err := fmt.Scanf("%d %d %d %d %d\n", &a, &b, &c, &d, &e)
	if err != nil {
		log.Fatalln(err)
	}

	s = append(s, a, b, c, d, e)
	for i, v := range s {
		s[i] = v * v
	}

	fmt.Println("The values of the slice have now been raised to the power of 2:")
	for _, v := range s {
		fmt.Printf("%d ", v)
	}

}
