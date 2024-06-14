package main

import (
	"fmt"
)

func main() {
	var number1, number2, listing, number, number_square int
	var eternity bool = false
	for eternity == false {
		listing = 1
		fmt.Println("Please choose two number")
		fmt.Scanln(&number1, &number2)
		switch {
		case number1 < number2:
			number = number1
			for number <= number2 {
				number_square = number * number
				fmt.Println(listing, ". number is", number, "and its square is ", number_square)
				number = number + 1
				listing = listing + 1
			}
			break
		case number2 < number1:
			number = number2
			for number <= number1 {
				number_square = number * number
				fmt.Println(listing, ". number is", number, "and its square is ", number_square)
				number = number + 1
				listing = listing + 1
			}
			break
		default:
			fmt.Println("Please choose two different number.")
		}
	}
}
