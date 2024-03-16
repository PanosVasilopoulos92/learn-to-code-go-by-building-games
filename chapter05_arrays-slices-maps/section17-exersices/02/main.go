package main

/*
Problem 02* - Simple calculator
Write a program which reads commands from the terminal & performs operations on a set of numbers.

The commands are:
add <num> - add a new number to the collection of numbers
inc <index> - increment a number at a given index (zero-based)
sub <index> <num> - subtract num from the number at the given index
mul <index> <num> - multiply the number at the given index by num
acc <index> <num> - accumulate num to the number at the given index
show - print the current value of all numbers at a single line
exit - stop program execution

If the provided index is outside the range of numbers you currently have, print “invalid index”.
If the operation was executed successfully, print “OK”.

*/

import (
	"fmt"
)

func main() {
	xi := []int{}
	fmt.Println("Calculator program has started...")
	displayMenu()

	for {
		var command string
		index := 0
		num := 0

		fmt.Println("Provide the command first:")
		_, err := fmt.Scanf("%s\n", &command)
		if err != nil {
			fmt.Scanln(&command)
			fmt.Println("Error while reading input. Try again.", err)
			continue
		}

		if command == "exit" {
			break
		}

		switch command {
		case "add":
			fmt.Println("Now provide the number you want to add to slice:")
			_, err := fmt.Scanf("%d\n", &num)
			if err != nil {
				fmt.Scanln(&command)
				fmt.Println("Error while reading input. Try again.", err)
				break
			}
			xi = append(xi, num)
			fmt.Printf("%d was successfully added to slice.\n", num)
		case "inc", "sub", "mul", "acc":
			fmt.Println("Now provide the index and the number you want:")
			_, err := fmt.Scanf("%d %d\n", &index, &num)
			if err != nil {
				fmt.Scanln(&index, &num)
				fmt.Println("Error while reading input. Try again.", err)
				break
			}
			if index < 0 || index >= len(xi) {
				fmt.Println("wrong index")
				break
			}

			switch command {
			case "inc":
				xi[index] += num
				fmt.Printf("Value at index %d successfully incremented by %d\n", index, num)
			case "sub":
				xi[index] = xi[index] - num
				fmt.Printf("Value at index %d is now %d\n", index, xi[index])
			case "mul":
				xi[index] = xi[index] * num
				fmt.Printf("Value at index %d is now %d\n", index, xi[index])
			case "acc":
				xi[index] = xi[index] + num
				fmt.Printf("Value at index %d is now %d\n", index, xi[index])
			}

		case "show":
			for _, v := range xi {
				fmt.Printf("%d ", v)
			}
			fmt.Println()
		case "showCommands":
			displayMenu()
		default:
			fmt.Println("wrong command. Try again.")
		}
	}
}

func displayMenu() {
	fmt.Println("The available commands are:")
	fmt.Println(`add <num> - add a new number to the collection of numbers
inc <index> - increment a number by 1 at a given index (zero-based)
sub <index> <num> - subtract num from the number at the given index
mul <index> <num> - multiply the number at the given index by num
acc <index> <num> - accumulate num to the number at the given index
show - print the current value of all numbers at a single line
show commands - displays the available commands
exit - stop program execution`)
}
