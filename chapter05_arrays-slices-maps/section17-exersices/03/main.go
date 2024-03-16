package main

/*
Problem 03* - Sort an array using Selection Sort
Learn about Selection Sort & sort an input array of numbers using the algorithm.

*/

import (
	"fmt"
	"log"
)

func main() {
	xi := []int{}
	var a, b, c, d, e, f, g, h int
	fmt.Println("Provide me with 8 integer numbers, in order to sort them:")
	_, err := fmt.Scanf("%d %d %d %d %d %d %d %d\n", &a, &b, &c, &d, &e, &f, &g, &h)
	if err != nil {
		log.Fatalln(err)
	}

	xi = append(xi, a, b, c, d, e, f, g, h)

	for i := 0; i <= len(xi)-2; i++ {
		minIndex := i
		for j := i + 1; j < len(xi); j++ {
			if xi[j] < xi[minIndex] {
				minIndex = j
			}
		}
		xi[i], xi[minIndex] = xi[minIndex], xi[i]
	}

	for v := range xi {
		fmt.Printf("%d ", v)
	}

}
