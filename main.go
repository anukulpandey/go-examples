package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now().Year()
	fmt.Println("year:", t)

	switch t {
	case 2023:
		fmt.Println("year is 2023")
	default:
		fmt.Println("invalid year")
	}
}
