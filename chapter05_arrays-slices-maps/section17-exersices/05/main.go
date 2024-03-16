package main

/*
Problem 05 - Excluded words
Write a program which takes in two lines of words and excludes all words in the second list that exist in the first one.
Also all values in the final list must be unique.

*/

import (
	"fmt"
	"log"
	"slices"
)

func main() {
	xs1 := make([]string, 0)
	xs2 := make([]string, 0)
	xs3 := make([]string, 0)
	var w1, w2, w3 string
	var w4, w5, w6 string
	fmt.Println("Provide 3 words:")
	_, err := fmt.Scanf("%s %s %s\n", &w1, &w2, &w3)
	if err != nil {
		log.Fatalln(err)
	}
	xs1 = append(xs1, w1, w2, w3)

	fmt.Println("Now provide the next 3 words:")
	_, err = fmt.Scanf("%s %s %s\n", &w4, &w5, &w6)
	if err != nil {
		log.Fatalln(err)
	}
	xs2 = append(xs2, w4, w5, w6)

	for i := 0; i < len(xs1); i++ {
		if !slices.Contains(xs3, xs1[i]) {
			xs3 = append(xs3, xs1[i])
		}
		for j := 0; j < len(xs2); j++ {
			if !slices.Contains(xs3, xs2[j]) {
				xs3 = append(xs3, xs2[j])
			}
		}
	}

	fmt.Println(xs3)

}
