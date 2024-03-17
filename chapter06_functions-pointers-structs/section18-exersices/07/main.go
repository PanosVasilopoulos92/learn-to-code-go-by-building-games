package main

import (
	"fmt"
)

/*
Problem 07 - Mutating points
Create three functions - incPtr, incNoPtr and incNewPtr which all take in a point and increment its values by 1. Here’s the catch:
incPtr should take in a pointer to a point & modify its contents
incNoPtr should take a point without a pointer & modify its contents
incNewPtr should take in a pointer and assign it a new Point with the incremented values

Print the values of the three points after being incremented. Anything unusual?

*/

type Point struct {
	row int
	col int
}

func main() {
	p1 := &Point{10, 20}
	p2 := Point{11, 21}
	p3 := &Point{12, 22}
	fmt.Println(p1)
	fmt.Printf("%T\n", p1)
	fmt.Println(p2)
	fmt.Printf("%T\n", p2)
	fmt.Println(p3)
	fmt.Printf("%T\n", p3)
	fmt.Println("--------")

	incPtr(p1)
	fmt.Println(p1)

	incNoPtr(p2)
	fmt.Println(p2)

	incNewPtr(p3)
	fmt.Println(p3)
}

func incPtr(p *Point) {
	p.row += 1
	p.col += 1
	fmt.Printf("%T\n", p)
}

func incNoPtr(p Point) {
	p.row++
	p.col++
	fmt.Printf("%T\n", p)
}

func incNewPtr(p *Point) {
	p2 := &Point{row: p.row + 1, col: p.col + 1}
	p = p2
	fmt.Printf("%T\n", p2)
}
