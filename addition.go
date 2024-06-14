package main

import (
	"fmt"
)

func main() {
	var number, result int
	var eternity bool = false

	for eternity == false {
		fmt.Println("Write to me a number.")
		fmt.Scanln(&number)
		if number > 0 {
			result = (number * (number + 1)) / 2
			fmt.Println("1 to ", number, "adition is ", result)
		} else {
			fmt.Println("Prgogram is closing, please wait")
			eternity = true
		}
	}
}
