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
		if number1 != number2 {
			for number1 <= number2 {
				fmt.Println(number1 * number1)
				number1 += 1
			}
			for number1 >= number2 {
				fmt.Println(number2 * number2)
				number2 += 1
			}
		} else {
			eternity = false
		}
	}
}
