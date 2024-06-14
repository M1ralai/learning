package main

import (
	"fmt"
)

func main() {
	var number1, number2, number3 int
	var eternity bool = false
	for eternity == false {
		fmt.Println("Please give me three number.")
		fmt.Scanln(&number1, &number2, &number3)
		fmt.Println("First number is: ", number1)
		fmt.Println("Second number is: ", number2)
		fmt.Println("Third number is: ", number3)
		if number1 == 0 {
			fmt.Println("Program has ended thank you to choosed us")
			eternity = true
		}
	}
}
