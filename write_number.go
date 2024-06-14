package main

import (
	"fmt"
)

func main() {
	var number, start int
	var eternity bool = false
	for eternity == false {
		start = 1
		fmt.Println("I will write until ")
		fmt.Scanln(&number)
		if number > 0 {
			fmt.Println("Here is your serie")
			for start <= number {
				fmt.Println(start)
				start = start + 1
			}
		} else {
			eternity = true
			fmt.Println("This program is closing, please wait.")
		}
	}
}
