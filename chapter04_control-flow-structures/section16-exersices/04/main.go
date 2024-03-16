package main

import (
	"fmt"
	"time"
)

/*
Problem 04 - Game shopping buddy
Write a program which buys the best possible choice, among three games you love, based on an input number, which represents the cash your parents have given you.
There are three games you love, in this order:

StarCraft 2 - 30$
Cyberpunk 2077 - 130$
The Witcher 3 - 60$

Given an input budget N, decide which games you’re going to buy. If you can afford all of them, then have them all!

If you can afford only 2 out of three, then make your preference in this order based on the budget you have.

The output should be in the following format:
Here’s what I bought:
{game1}, {game2}, {game 3}

If you bought less than three games, print the games you did buy in this order.
If you couldn’t buy anything, print I couldn’t buy anything!

*/

func main() {
	const (
		starCraft_2    = 30.0
		cyberpunk_2077 = 130.0
		the_Witcher_3  = 60.0
	)
	var totalMoney float64
	var remainingMoney float64
	canIBuy := true

	fmt.Println("Provide the amount of money that your parents gave to you:")
	fmt.Scan(&totalMoney)
	fmt.Printf("Ok, you have %.2f in total. Let's see what games you can have...\n", totalMoney)
	time.Sleep(2 * time.Second)

	switch canIBuy {
	case totalMoney >= starCraft_2+cyberpunk_2077+the_Witcher_3:
		remainingMoney = totalMoney - (starCraft_2 + cyberpunk_2077 + the_Witcher_3)
		fmt.Printf("Excellent, you can have all of them. Your current outstand is now: %.2f", remainingMoney)
	case totalMoney >= starCraft_2+cyberpunk_2077:
		remainingMoney = totalMoney - (starCraft_2 + cyberpunk_2077)
		fmt.Printf("You just bought StarCraft 2 and Cyberpunk 2077. Your current outstand is now: %.2f", remainingMoney)
	case totalMoney >= starCraft_2+the_Witcher_3:
		remainingMoney = totalMoney - (starCraft_2 + the_Witcher_3)
		fmt.Printf("You just bought StarCraft 2 and The Witcher 3. Your current outstand is now: %.2f", remainingMoney)
	case totalMoney >= cyberpunk_2077:
		remainingMoney = totalMoney - cyberpunk_2077
		fmt.Printf("You just bought Cyberpunk 2077. Your current outstand is now: %.2f", remainingMoney)
	case totalMoney >= the_Witcher_3:
		remainingMoney = totalMoney - the_Witcher_3
		fmt.Printf("You just bought The Witcher 3. Your current outstand is now: %.2f", remainingMoney)
	case totalMoney >= starCraft_2:
		remainingMoney = totalMoney - starCraft_2
		fmt.Printf("You just bought StarCraft 2. Your current outstand is now: %.2f", remainingMoney)
	default:
		fmt.Printf("Sorry, you cannot buy anything.")
	}

}
