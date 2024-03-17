package main

/*
Look at file 'requirements' for implementation.
*/

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

type Gopher struct {
	name      string
	hitpoints int
	weapon    Weapon
	strength  int
	agility   int
	intellect int
	coins     int
}

type Weapon struct {
	name            string
	damageDeal      []int
	price           float64
	strengthReq     int
	agilityReq      int
	intelligenceReq int
}

type Consumable struct {
	name            string
	hitpointsEffect int
}

var actions = []string{"Attack", "BuyWeapon", "BuyConsumable", "Work", "TrainSkill", "Retreat"}
var gophers = createPlayers()
var weapons = createWeapons()
var healthPotions = createConsumables()
var gopherPlaying, opponent = 0, 1
var playerSurrendered = false
var successfulPersuade bool

func main() {
	numOfTurns := 0
	var user1, user2 string

	fmt.Println("Please provide your names separated by one space:")
	_, err := fmt.Scanf("%s %s \n", &user1, &user2)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println()
	gophers[0].name = user1
	gophers[1].name = user2

	fmt.Printf("So, let the game begin!! %s will start first.\n", gophers[gopherPlaying].name)
	fmt.Println()

	for exitGame := false; !exitGame; {
		numOfTurns++
		fmt.Printf("Round %d started.\n", numOfTurns)

		if numOfTurns%2 != 0 {
			gopherPlaying = 0
			opponent = 1
		} else {
			gopherPlaying = 1
			opponent = 0
		}

		makeAction()

		if gophers[opponent].hitpoints <= 0 {
			fmt.Printf("%s is the winner!", gophers[gopherPlaying].name)
			exitGame = true
		} else if playerSurrendered {
			fmt.Printf("%s is the winner!", gophers[opponent].name)
			exitGame = true
		} else {
			fmt.Printf("End of turn %d.\t", numOfTurns)
		}
	}
}

func chooseAction() int {
	chosenAction := 0

	fmt.Println("It's your turn dear gopher. Here is your current stats:")
	printCurrentStats()
	fmt.Printf("%s, choose one of the following actions by just inserting the corresponding number.\nChoose wisely...\n", gophers[gopherPlaying].name)
	for i, v := range actions {
		fmt.Printf("%d. %s\n", i+1, v)
	}
	for {
		_, err := fmt.Scanf("%d\n", &chosenAction)
		if err != nil {
			log.Fatalln(err)
			continue
		}
		return chosenAction
	}
}

func makeAction() {
	actionMade := chooseAction()

	switch actionMade {
	case 1:
		attack(&gophers[gopherPlaying], &gophers[opponent])
		time.Sleep(1 * time.Second)
		fmt.Println()
	case 2:
		weaponPersuaded := ""
		for {
			fmt.Println("Ok, type the name of weapon you want to buy. Available weapons are:")
			for _, v := range weapons[1:] {
				fmt.Printf("%s ", v.name)
			}
			fmt.Println()

			_, err := fmt.Scanf("%s\n", &weaponPersuaded)
			if err != nil {
				fmt.Println("Not such a weapon exist. Try again:")
			}
			if !buyWeapon(&gophers[gopherPlaying], weaponPersuaded) {
				continue
			} else {
				fmt.Printf("Great. You now have %s as a weapon.\n", gophers[gopherPlaying].weapon.name)
				fmt.Println()
				time.Sleep(2 * time.Second)
				break
			}
		}

	case 3:
		potionPersuaded := ""
		for {
			fmt.Println("Ok, type the the potion you want to buy. Available potions are:")
			for i, v := range healthPotions {
				fmt.Printf("%d.%s, which gives %d health\n", i+1, v.name, v.hitpointsEffect)
			}
			fmt.Println()

			_, err := fmt.Scanf("%s\n", &potionPersuaded)
			if err != nil {
				fmt.Println("Not such a potion exist. Try again:")
			}

			if buyPotion(&gophers[gopherPlaying], potionPersuaded) != "successful transaction" {
				fmt.Println("Error in transaction. You have to try again.")
				continue
			}
			fmt.Printf("Your health is now %d.\n", gophers[gopherPlaying].hitpoints)
			fmt.Println()
			time.Sleep(2 * time.Second)
			break
		}
	case 4:
		work(&gophers[gopherPlaying])
		fmt.Println()
	case 5:
		skill := ""
		for {
			fmt.Println("Ok, type the name of skill you want to upgrade.\nAvailable skills are: strength, agility and intellect.")
			_, err := fmt.Scanf("%s\n", &skill)
			if err != nil {
				fmt.Println("error in input")
			}
			if !trainSkill(&gophers[gopherPlaying], skill) {
				fmt.Println("Not such a skill exist. Try again:")
				continue
			} else {
				time.Sleep(2 * time.Second)
				fmt.Println()
				break
			}
		}
	case 6:
		for {
			response := surrender()
			if response == "Yes" {
				fmt.Println("You can now live in shame...")
				playerSurrendered = true
				os.Exit(1)
			} else if response == "No" {
				fmt.Println("Wise choice friend. We, gophers, never surrender!!")
				fmt.Println("... But actions have always consequences. You lose your turn.")
				fmt.Println()
				time.Sleep(2 * time.Second)
				break
			} else {
				fmt.Println("Not a valid choice. Answer again.")
				continue
			}
		}
	}
}

func attack(player, opponent *Gopher) {

	// Creates a new random number generator
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	damage := 0

	if player.weapon.name == weapons[0].name {
		damage = 1
		opponent.hitpoints -= damage
	}
	if player.weapon.name == weapons[1].name {
		damage = random.Intn(2) + 2
		opponent.hitpoints -= damage
	}
	if player.weapon.name == weapons[2].name {
		damage = random.Intn(3) + 3
		opponent.hitpoints -= damage
	}
	if player.weapon.name == weapons[3].name {
		damage = random.Intn(7) + 1
		opponent.hitpoints -= damage
	}
	if player.weapon.name == weapons[4].name {
		damage = 3
		opponent.hitpoints -= damage
	}
	if player.weapon.name == weapons[5].name {
		damage = random.Intn(3)*3 + 1
		opponent.hitpoints -= damage
	}
	if player.weapon.name == weapons[6].name {
		damage = random.Intn(2) + 6
		opponent.hitpoints -= damage
	}
	if player.weapon.name == weapons[7].name {
		damage = 4
		opponent.hitpoints -= damage
	}

	fmt.Printf("Attack was successful! %s lost %d HP", opponent.name, damage)
}

func buyWeapon(gopher *Gopher, weapon string) bool {
	successfulPersuade = false
	msgNotEnoughMoney := "Sorry, you cannot afford this weapon. You can choose another one."

	switch weapon {
	case "knife":
		if gopher.coins >= 10 {
			gopher.weapon = weapons[1]
			successfulPersuade = true
			gopher.coins -= 10
		} else {
			fmt.Println(msgNotEnoughMoney)
		}
	case "sword":
		if gopher.coins >= 35 {
			gopher.weapon = weapons[2]
			successfulPersuade = true
			gopher.coins -= 35
		} else {
			fmt.Println(msgNotEnoughMoney)
		}
	case "ninjaku":
		if gopher.coins >= 25 {
			gopher.weapon = weapons[3]
			successfulPersuade = true
			gopher.coins -= 25
		} else {
			fmt.Println(msgNotEnoughMoney)
		}
	case "wand":
		if gopher.coins >= 30 {
			gopher.weapon = weapons[4]
			successfulPersuade = true
			gopher.coins -= 30
		} else {
			fmt.Println(msgNotEnoughMoney)
		}
	case "gophermourne":
		if gopher.coins >= 65 {
			gopher.weapon = weapons[5]
			successfulPersuade = true
			gopher.coins -= 65
		} else {
			fmt.Println(msgNotEnoughMoney)
		}
	case "warglaivesOfGopherinoth":
		if gopher.coins >= 55 {
			gopher.weapon = weapons[6]
			successfulPersuade = true
			gopher.coins -= 55
		} else {
			fmt.Println(msgNotEnoughMoney)
		}
	case "codeseeker":
		if gopher.coins >= 60 {
			gopher.weapon = weapons[7]
			successfulPersuade = true
			gopher.coins -= 60
		} else {
			fmt.Println(msgNotEnoughMoney)
		}
	default:
		successfulPersuade = false
	}
	return successfulPersuade
}

func buyPotion(gopher *Gopher, potion string) string {
	switch potion {
	case "smallPotion":
		if gopher.coins >= 5.0 {
			gopher.hitpoints += 5
			gopher.coins -= 5.0
			fmt.Print("Great, you gained 5 health. ")
		} else {
			fmt.Println("Sorry, not enough money...")
			return "unsuccessful transaction"
		}
	case "mediumPotion":
		if gopher.coins >= 9.0 {
			gopher.hitpoints += 10
			gopher.coins -= 9.0
			fmt.Print("Great, you gained 10 health. ")
		} else {
			fmt.Println("Sorry, not enough money...")
			return "unsuccessful transaction"
		}
	case "bigPotion":
		if gopher.coins >= 18.0 {
			gopher.hitpoints += 20
			gopher.coins -= 18.0
			fmt.Print("Super! You gained 20 health. ")
		} else {
			fmt.Println("Sorry, not enough money...")
			return "unsuccessful transaction"
		}
	default:
		return "unsuccessful transaction"
	}
	return "successful transaction"
}

func work(gopher *Gopher) {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	coinsEarned := random.Intn(11) + 5

	gopher.coins += coinsEarned
	fmt.Printf("Great, you earned %d coins from your work. You now have %d coins\n", coinsEarned, gopher.coins)
	time.Sleep(2 * time.Second)
}

func trainSkill(gopher *Gopher, skill string) bool {
	updateSuccess := false
	switch skill {
	case "strength":
		if gopher.coins >= 5 {
			gopher.strength += 2
			gopher.coins -= 5
			fmt.Printf("You gained 2 strength points. Your strength is now level %d.\n", gopher.strength)
			updateSuccess = true
		}
	case "agility":
		if gopher.coins >= 5 {
			gopher.agility += 2
			gopher.coins -= 5
			fmt.Printf("You gained 2 agility points. Your agility is now level %d.\n", gopher.agility)
			updateSuccess = true
		}
	case "intellect":
		if gopher.coins >= 5 {
			gopher.intellect += 2
			gopher.coins -= 5
			fmt.Printf("You gained 2 intellect points.Your intellect is now level %d.\n", gopher.intellect)
			updateSuccess = true
		}
	default:
		updateSuccess = false
	}
	return updateSuccess
}

func surrender() string {
	response := ""
	for {
		fmt.Println("Are you really want to surrender? Type 'Yes' or 'No' gopher.")
		_, err := fmt.Scanf("%s\n", &response)
		if err != nil {
			fmt.Println("Not a valid answer. Type again.")
			continue
		}
		break
	}
	return response
}

func createPlayers() []Gopher {
	xg := []Gopher{}

	gopher1 := Gopher{
		hitpoints: 30,
		weapon:    weapons[0],
		coins:     20,
	}
	gopher2 := Gopher{
		hitpoints: 30,
		weapon:    weapons[0],
		coins:     20,
	}

	xg = append(xg, gopher1, gopher2)
	return xg

}

func createConsumables() []Consumable {
	xs := []Consumable{}

	smallPotion := Consumable{
		name:            "smallPotion",
		hitpointsEffect: 5,
	}
	mediumPotion := Consumable{
		name:            "mediumPotion",
		hitpointsEffect: 10,
	}

	bigPotion := Consumable{
		name:            "bigPotion",
		hitpointsEffect: 20,
	}

	xs = append(xs, smallPotion, mediumPotion, bigPotion)
	return xs
}

func createWeapons() []Weapon {
	xw := []Weapon{}

	bareHands := Weapon{
		name:       "bareHands",
		damageDeal: []int{1},
	}
	knife := Weapon{
		name:       "knife",
		damageDeal: []int{2, 3},
		price:      10.0,
	}
	sword := Weapon{
		name:        "sword",
		damageDeal:  []int{3, 5},
		price:       35.0,
		strengthReq: 2,
	}
	ninjaku := Weapon{
		name:       "ninjaku",
		damageDeal: []int{1, 7},
		price:      25.0,
		agilityReq: 2,
	}
	wand := Weapon{
		name:            "wand",
		damageDeal:      []int{3, 3},
		price:           30.0,
		intelligenceReq: 2,
	}
	gophermourne := Weapon{
		name:        "gophermourne",
		damageDeal:  []int{6, 7},
		price:       65.0,
		strengthReq: 5,
	}
	warglaivesOfGopherinoth := Weapon{
		name:       "warglaivesOfGopherinoth",
		damageDeal: []int{2, 9},
		price:      55.0,
		agilityReq: 5,
	}
	codeseeker := Weapon{
		name:            "codeseeker",
		damageDeal:      []int{4, 4},
		price:           60.0,
		intelligenceReq: 5,
	}

	xw = append(xw, bareHands, knife, sword, ninjaku, wand, gophermourne, warglaivesOfGopherinoth, codeseeker)
	return xw
}

func printCurrentStats() {
	fmt.Printf("HP: %d\n", gophers[gopherPlaying].hitpoints)
	fmt.Printf("Weapon: %s\n", gophers[gopherPlaying].weapon.name)
	fmt.Printf("Coins: %d\n", gophers[gopherPlaying].coins)
	fmt.Printf("Strength: %d\n", gophers[gopherPlaying].strength)
	fmt.Printf("Agility: %d\n", gophers[gopherPlaying].agility)
	fmt.Printf("Intellect: %d\n", gophers[gopherPlaying].intellect)
}
