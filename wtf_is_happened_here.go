package main

import (
	"fmt"
)

func main() {
	var n1, n2 int
	fmt.Println("Please give me 2 number.")
	fmt.Scanln(&n1, &n2)
	fmt.Println("Your numbers are", n1, "and ", n2)
	operation(n1, n2)

}
func operation(n1, n2 int) {
	if n1 == n2 {
		println("Please write different numbers.")
		main()
	} else if n1 >= n2 {
		for n1 >= n2 {
			fmt.Println(n2 * n2)
			n2 += 1
		}
	} else {
		operation(n2, n1)
	}

}
