package main

/*
Problem 06 - PrintAt with Point
Modify the PrintAt function you created in Problem 02 to accept a Point instead of a row and col.
The function should still work as expected.

*/

import (
	"time"

	tm "github.com/buger/goterm"
)

type Point struct {
	row int
	col int
	msg time.Time
}

func main() {
	p1 := &Point{4, 10, time.Now()}
	printAt(p1)
	time.Sleep(2 * time.Second)

	p2 := &Point{2, 10, time.Now()}
	printAt(p2)
	time.Sleep(2 * time.Second)
}

func printAt(p *Point) {
	tm.Clear()
	tm.MoveCursor(p.col, p.row)
	tm.Print(tm.Color(p.msg.Local().String(), tm.GREEN))
	tm.Flush()
}
