package main

/*
Look at file 'requirements' for implementation.
*/

import (
	"fmt"
	"math/rand"
	"time"
)

type Gopher struct {
	hitpoints int
	weapon    Weapon
	strength  int
	agility   int
	intellect int
	coins     int
}

type Weapon struct {
	damageDeal      []int
	price           float64
	strengthReq     int
	agilityReq      int
	intelligenceReq int
}

type Consumable struct {
	hitpointsEffect int
}

var actions = []string{"Attack", "BuyWeapon", "BuyConsumable", "Work", "TrainSkill", "Retreat"}

func main() {
	gophers := createPlayers()
	weapons := createWeapons()
	healthPotions := createConsumables()
	numOfTurns := 0

	fmt.Println("Let the game begin!! gopher1 will start first. Fight!")

	for exitGame := false; !exitGame; {
		gopherPlaying := 0
		numOfTurns++
		makeAction()
	}

}

func chooseAction() int {
	chosenAction := 0
	fmt.Println("It's your turn dear gopher. Choose one of the following actions by just inserting\nthe corresponding number. Choose wisely...")
	for i, v := range actions {
		fmt.Printf("%d. %s\n", i, v)
	}
	for {
		_, err := fmt.Scanf("%d\n", &chosenAction)
		if err != nil {
			fmt.Println("Wrong input my friend, choose again more carefully:")
			continue
		}
		return chosenAction
	}
}

func makeAction() {
	gophers := createPlayers()
	weapons := createWeapons()
	actionMade := chooseAction()
	switch actionMade {
	case 1:
		attack(&gophers[0], &gophers[1])
	case 2:
		weaponPersuaded := ""
		for {
			fmt.Println("Ok, type the name of weapon you want to buy. Available weapons are:")
			for _, v := range weapons {
				fmt.Printf("%s ", v)
			}
			_, err := fmt.Scanf("%s\n", &weaponPersuaded)
			if err != nil {
				fmt.Println("Not such a weapon. Try again:")
				continue
			}
			break
		}
		buyWeapon(&gophers[0], weaponPersuaded)
	case 3:
	}
}

func attack(gopher *Gopher, gophers []Gopher) {
	weapons := createWeapons()
	gophers = createPlayers()

	// Create a new random number generator
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	if &gopher1.weapon == &weapons[0] {
		gopher2.hitpoints -= 1

	}
	if &gopher1.weapon == &weapons[1] {
		damage := random.Intn(2) + 2
		gopher2.hitpoints -= damage
	}
	if &gopher1.weapon == &weapons[2] {
		damage := random.Intn(3) + 3
		gopher2.hitpoints -= damage
	}
	if &gopher1.weapon == &weapons[3] {
		damage := random.Intn(7) + 1
		gopher2.hitpoints -= damage
	}
	if &gopher1.weapon == &weapons[4] {
		gopher2.hitpoints -= 3
	}
	if &gopher1.weapon == &weapons[5] {
		damage := random.Intn(3)*3 + 1
		gopher2.hitpoints -= damage
	}
	if &gopher1.weapon == &weapons[6] {
		damage := random.Intn(2) + 6
		gopher2.hitpoints -= damage
	}
	if &gopher1.weapon == &weapons[7] {
		gopher2.hitpoints -= 4
	}
}

func buyWeapon(gopher *Gopher, weapon string) {
	weapons := createWeapons()
	switch weapon {
	case "knife":
		gopher.weapon = weapons[1]
	case "sword":
		gopher.weapon = weapons[2]
	case "ninjaku":
		gopher.weapon = weapons[3]
	case "wand":
		gopher.weapon = weapons[4]
	case "gophermourne":
		gopher.weapon = weapons[5]
	case "warglaivesOfGopherinoth":
		gopher.weapon = weapons[6]
	case "codeseeker":
		gopher.weapon = weapons[7]
	default:
		fmt.Println("...")
	}
}

func createPlayers() []Gopher {
	xg := []Gopher{}
	weapons := createWeapons()
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
	// consumables := make(map[string]float64)
	// consumables["small_health_potion"] = 5.0  // hitpointsEffect = 5
	// consumables["medium_health_potion"] = 9.0 // hitpointsEffect = 10
	// consumables["big_health_potion"] = 18.0   // hitpointsEffect = 20
	smallPotion := Consumable{
		hitpointsEffect: 5,
	}
	mediumPotion := Consumable{
		hitpointsEffect: 10,
	}

	bigPotion := Consumable{
		hitpointsEffect: 20,
	}

	xs = append(xs, smallPotion, mediumPotion, bigPotion)
	return xs
}

func createWeapons() []Weapon {
	xw := []Weapon{}

	bareHands := Weapon{
		damageDeal: []int{1},
	}
	knife := Weapon{
		damageDeal: []int{2, 3},
		price:      10.0,
	}
	sword := Weapon{
		damageDeal:  []int{3, 5},
		price:       35.0,
		strengthReq: 2,
	}
	ninjaku := Weapon{
		damageDeal: []int{1, 7},
		price:      25.0,
		agilityReq: 2,
	}
	wand := Weapon{
		damageDeal:      []int{3, 3},
		price:           30.0,
		intelligenceReq: 2,
	}
	gophermourne := Weapon{
		damageDeal:  []int{6, 7},
		price:       65.0,
		strengthReq: 5,
	}
	warglaivesOfGopherinoth := Weapon{
		damageDeal: []int{2, 9},
		price:      55.0,
		agilityReq: 5,
	}
	codeseeker := Weapon{
		damageDeal:      []int{4, 4},
		price:           60.0,
		intelligenceReq: 5,
	}

	xw = append(xw, bareHands, knife, sword, ninjaku, wand, gophermourne, warglaivesOfGopherinoth, codeseeker)
	return xw
}
