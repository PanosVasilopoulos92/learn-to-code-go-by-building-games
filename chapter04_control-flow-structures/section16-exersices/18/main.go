package main

import (
	"fmt"
	"log"
)

/*
Problem 18* - Shuffle cards
You are given a set of standard playing cards. They are represented by the numbers 2 through 10 and the characters A, J, Q, K.

Write a program to shuffle the cards in a random order.

*/

func main() {
	var s = []string{}
	var a, b, c, d, e string
	fmt.Println("Choose 5 cards between 2-10 and A, J, Q, K. Values must be separated by space between:")
	_, err := fmt.Scanf("%s %s %s %s %s\n", &a, &b, &c, &d, &e)
	if err != nil {
		fmt.Scan(&s)
		log.Fatalln(err)
	}

	s = append(s, a, b, c, d, e)
	for _, v := range s {
		fmt.Println(v)
	}

}
