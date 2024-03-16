package main

import (
	"fmt"
	"log"
	"time"
)

/*
Problem 08* - Gaming time
Write a program which checks if it’s time to play on your gaming PC.

You’re given an input time in the format {hour}:{minutes}:{seconds} AM/PM and you should check if the time is in one of the ranges when you can play on your PC:
07:00:00 PM - 09:00:00 PM
08:00:00 AM - 10:00:00 AM
11:00:00 PM - 06:00:00 AM

If the given time is in any of those ranges, print It’s gaming time!
If not, print It’s not the time for games yet…

*/

func main() {
	var hour, minutes, seconds int

	period1Start := time.Date(0, 0, 0, 19, 0, 0, 0, time.UTC)
	period1End := time.Date(0, 0, 0, 21, 0, 0, 0, time.UTC)
	period2Start := time.Date(0, 0, 0, 8, 0, 0, 0, time.UTC)
	period2End := time.Date(0, 0, 0, 10, 0, 0, 0, time.UTC)
	period3Start := time.Date(0, 0, 0, 23, 0, 0, 0, time.UTC)
	period3End := time.Date(0, 0, 0, 6, 0, 0, 0, time.UTC)

	fmt.Println(period1Start, period1End)
	fmt.Println(period2Start, period2End)
	fmt.Println(period3Start, period3End)

	fmt.Println("Provide the desired time in this format: {hour}:{minutes}:{seconds} AM/PM.")
	n, err := fmt.Scanf("%d:%d:%d\n", &hour, &minutes, &seconds)
	if err != nil {
		log.Fatalln(err)
	}
	if n != 3 {
		fmt.Println(n)
		log.Fatalln("Error in the number of input arguments.")
	}

	periodToCheck := time.Date(0, 0, 0, hour, minutes, seconds, 0, time.UTC)
	isInPeriod1 := periodToCheck.After(period1Start) && periodToCheck.Before(period1End)
	isInPeriod2 := periodToCheck.After(period2Start) && periodToCheck.Before(period2End)
	isInPeriod3 := periodToCheck.After(period3Start) && periodToCheck.Before(period3End)

	if isInPeriod1 || isInPeriod2 || isInPeriod3 {
		fmt.Printf("It's %v so... It’s gaming time!", periodToCheck)
	} else {
		fmt.Printf("It's %v so... It’s not the time for games yet…", periodToCheck)
	}

}
