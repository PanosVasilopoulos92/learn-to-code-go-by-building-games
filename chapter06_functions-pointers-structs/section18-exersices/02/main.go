package main

/*
Problem 02 - Create a PrintAt function
Using the 'buger/goterm' library, create a PrintAt function which takes in a row, col and string and prints the string at the designated position on the terminal.

*/

import (
	"time"

	tm "github.com/buger/goterm"
)

func main() {
	tm.Clear()
	printAt(5, 2, time.Now())

	time.Sleep(5 * time.Second)

}

func printAt(col, row int, msg time.Time) {
	tm.MoveCursor(col, row)

	tm.Print(tm.Color(msg.Local().String(), tm.GREEN))
	tm.Flush()
}
