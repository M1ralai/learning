package main

import (
	"fmt"
)

func main() {
	var number int
	var eternity bool = false
	for eternity == false {
		fmt.Println("Write a number: (0 to close proggram)")
		fmt.Scanln(&number)
		switch {
		case number == 0:
			fmt.Println("Program is closing, please wait")
			eternity = true
		case number < 0:
			fmt.Println("Please write positive number")
		default:
			if number%2 == 0 {
				fmt.Println("This number is idiom")
			} else {
				fmt.Println("This number is odd")
			}
		}
	}
}
