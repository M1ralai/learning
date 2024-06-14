package main

import (
	"fmt"
)

func main() {
	var number1, number2 int
	var eternity bool = true
	for eternity == true {
		fmt.Println("Choose 2 number")
		fmt.Scanln(&number1, &number2)
		if number1 <= number2 {
			for number1 <= number2 {
				fmt.Println(number1 * number1)
				number1++
			}
		} else if number1 == number2 {
			fmt.Println("Program closingiplease wait.")
			eternity = false
		} else {
			fmt.Println("Choose first number greater")
		}

	}
}
