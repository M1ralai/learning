package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	var decesion int
	fmt.Println("Choose your destiny. \n 1. 1500 hp / 1000 armour / 2500 damage \n 2. 1000 hp / 2500 armour / 1000 damage")
	fmt.Scanln(&decesion)
	if decesion == 1 {
		monsters(1500, 1000, 2500)
	} else if decesion == 2 {
		monsters(1000, 2500, 1000)
	} else {
		fmt.Printf(" \nChoose one of the options.\n")
		main()
	}
}

func monsters(health, armour, damage int) {
	var rng int
	fmt.Println("LOOK!!! there is three MONSTERSS!!!")
	rng = rand.IntN(3)
	if rng == 0 {
		fmt.Println(" \nThere is a Scarab!!! \n")
		scarab(health, armour, damage)
	} else if rng == 1 {
		fmt.Println(" \nThere is a Zombie!!! \n")
		main()
	} else {
		fmt.Println(" \nThere is a Skeleton!!! \n")
		main()
	}
}
func scarab(hp, a, dmg int) {
	var sdmg, rng, shp, sa, gdmg, tdmg int
	sa = 200
	shp = 1500
	sdmg = 200
	for hp >= 0 && shp >= 0 {
		rng = rand.IntN(41)
		tdmg = sdmg * (rng + 80) / 100 * (100000 - a) / 100000
		hp = hp - tdmg
		fmt.Println("\nYour current hp: ", hp)
		fmt.Println("Scarab gived ", tdmg, "damage.\n\n")
		rng = rand.IntN(41)
		gdmg = dmg * (rng + 80) / 100 * (100000 - sa) / 100000
		shp = shp - gdmg
		fmt.Println("\n Scarab's health: ", shp)
		fmt.Println("You gived ", gdmg, "damage.\n\n")
	}
	if hp < 0 {
		fmt.Println("You Died")
	} else {
		fmt.Println("You beat that bastard!!!")
	}
}
