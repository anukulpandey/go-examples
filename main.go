package main

import (
	"fmt"
)

func main() {
	// basic loop
	i := 0
	for i < 3 {
		fmt.Println(i)
		i++
	}
	// conventional for loop
	for j := 0; j < 3; j += 1 {
		fmt.Println(j)
	}
	// for without condition
	for {
		fmt.Println("anukul")
		break
	}
	// continue
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue //skips the next code blocks
		}
		fmt.Println(n)
	}
}
