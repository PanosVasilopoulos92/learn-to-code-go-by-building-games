package main

/*
Problem 04 - Union of numbers
Write a program which takes in two lines of numbers and outputs the union of the two lists. There should be no duplicates.

Note: The order of the numbers in the final output doesn’t matter

*/

import (
	"fmt"
	"log"
	"slices"
)

func main() {
	var a, b, c, d, e int
	var f, g, h, i, j int
	xi := []int{}
	xj := []int{}
	xij := []int{}
	fmt.Println("Provide the first group of integer numbers:")
	_, err := fmt.Scanf("%d %d %d %d %d\n", &a, &b, &c, &d, &e)
	if err != nil {
		log.Fatalln(err)
	}
	xi = append(xi, a, b, c, d, e)

	fmt.Println("Provide the second group of integer numbers:")
	_, err = fmt.Scanf("%d %d %d %d %d\n", &f, &g, &h, &i, &j)
	if err != nil {
		log.Fatalln(err)
	}
	xi = append(xi, f, g, h, i, j)

	for i := 0; i < len(xi); i++ {
		if !slices.Contains(xij, xi[i]) {
			xij = append(xij, xi[i])
		}
		for j := 0; j < len(xj); j++ {
			if !slices.Contains(xij, xj[j]) {
				xij = append(xij, xj[j])
			}
		}
	}

	fmt.Println("Process completed. The union of the two provided slices is:")
	for _, v := range xij {
		fmt.Printf("%d ", v)
	}
}
