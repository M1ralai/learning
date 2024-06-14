package main

import (
	"fmt"
)

func main() {
	var mainnumber, modulenumber, module int
	var eternity bool = false
	for eternity == false {
		fmt.Print("Please give me a number: ")
		fmt.Scanln(&mainnumber)
		fmt.Print("Please give me module number: ")
		fmt.Scan(&modulenumber)
		switch {
		case mainnumber > modulenumber:
			module = mainnumber % modulenumber
			fmt.Println(mainnumber, " mod ", modulenumber, "is: ", module)
		case mainnumber == 0:
			eternity = true
			fmt.Println("Program is closing, please wait.")
		default:
			fmt.Println("Error aquired")
		}
	}
}
