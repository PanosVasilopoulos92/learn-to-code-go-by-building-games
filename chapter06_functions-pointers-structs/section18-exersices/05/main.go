package main

/*
Problem 05 - Point
Create a struct called Point which has two fields: row and col.

Create several instances of Point and print them on the terminal.

*/

import "fmt"

type Point struct {
	row int
	col int
}

func main() {
	p1 := Point{
		row: 2,
		col: 4,
	}

	p2 := Point{
		row: 10,
		col: 20,
	}

	p3 := Point{30, 40}

	p4 := Point{50, 100}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)
	fmt.Println(p4.col, p4.row)
}
