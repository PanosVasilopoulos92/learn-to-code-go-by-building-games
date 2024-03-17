package main

/*
Problem 01 - Game Instructions printer
Create a program which has a function printGameInstructions, which doesn’t take any parameters and doesn’t return a result.
It should print the instructions for the hangman game.

Instructions:
Have the player select a letter of the alphabet.
If the letter is contained in the word/phrase, all the the letters equal to it are revealed.
If the letter is not contained in the word/phrase, a portion of the hangman is added.
The game continues until:
the word/phrase is guessed (all letters are revealed) – WINNER or,
all the parts of the hangman are displayed – LOSER

*/

import "fmt"

func main() {
	printGameInstructions()

	fmt.Println(`
	|---------|
	|	  O
	|	  
	|	 
	|
	|	`)
	fmt.Println(`
	|---------|
	|	  O
	|	  |
	|	 
	|
	|	`)
	fmt.Println(`
	|---------|
	|	  O
	|	 /|
	|	 
	|
	|	`)
	fmt.Println(`
	|---------|
	|	  O
	|	 /|\
	|	 
	|
	|	`)
	fmt.Println(`
	|---------|
	|	  O
	|	 /|\
	|	 / 
	|
	|	`)
	fmt.Println(`
	|---------|
	|	  O
	|	 /|\
	|	 / \
	|
	|	`)
}

func printGameInstructions() {
	fmt.Print(`
1. Have the player select a letter of the alphabet.
2. If the letter is contained in the word/phrase, all the the letters equal to it are revealed.
3. If the letter is not contained in the word/phrase, a portion of the hangman is added.
4. The game continues until:
   a. the word/phrase is guessed (all letters are revealed) – WINNER or,
   b. all the parts of the hangman are displayed – LOSER...
`)
}
