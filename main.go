package main

import (
	"fmt"
)

func main() {
	const condition = 9 % 2
	if condition == 0 {
		fmt.Println("divisible")
	} else {
		fmt.Println("not divisible")
	}

	// in single line
	// if statement ; condition
	if condition := 9 % 2; condition == 0 {
		fmt.Println("divisible")
	} else {
		fmt.Println("not divisible")
	}

}
