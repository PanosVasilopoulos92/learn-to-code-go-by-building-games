package main

import (
	"fmt"
	"log"
)

/*
Problem 06 - Which game is this character from?
Write a program which determines which game a character belongs to from an input character name.

The characters which will be given by game are:
For The Witcher 3:
Geralt
Ciri
Yennefer


For StarCraft 2:
Aldaris
Jim Raynor
Kerrigan
Zeratul

For WarCraft 3:
Arthas
Illidan
Sylvanas

Based on the input name, print {character name} is part of the game {game name}!

If you get a character which is not in this list, print I don’t recognize this character…

*/

func main() {
	var nameOfCharacter string

	fmt.Println("Insert a character name in order to learn if this character belong to 'The Witcher 3', 'StarCraft 2' or 'WarCraft 3':")
	_, err := fmt.Scan(&nameOfCharacter)
	if err != nil {
		log.Fatalln("error while reading input. Program exits...")
	}

	switch nameOfCharacter {
	case "Geralt":
		fallthrough
	case "Ciri":
		fallthrough
	case "Yennefer":
		fmt.Printf("Character %s belongs to game 'The Witcher 3'.", nameOfCharacter)
	case "Aldaris":
		fallthrough
	case "Jim Raynor":
		fallthrough
	case "Kerrigan":
		fallthrough
	case "Zeratul":
		fmt.Printf("Character %s belongs to game 'StarCraft 2'.", nameOfCharacter)
	case "Arthas":
		fallthrough
	case "Illidan":
		fallthrough
	case "Sylvanas":
		fmt.Printf("Character %s belongs to game 'WarCraft 3'.", nameOfCharacter)
	default:
		fmt.Println("I don’t recognize this character...")
	}

}
