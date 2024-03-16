package main

/*
Problem 07 - Find the longest word
Write a program which finds the longest word in a given text.
If multiple longest words exist, print the leftmost one.
string for input: A whisper of wind through the trees, sunlight dappling fallen leaves.

*/

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var s string
	var lengthiestWord string
	fmt.Println("Provide a string and I will print the longest word that it contains:")
	reader := bufio.NewReader(os.Stdin)
	s, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln(err)
	}

	xs := strings.SplitAfter(s, " ")
	for _, v := range xs {
		tmp := v
		if len(v) > len(lengthiestWord) {
			lengthiestWord = tmp
		}
	}
	fmt.Println(lengthiestWord)

}
