package main

import (
	"fmt"
	"math"
)

func main() {
	var a = math.Max(12.3, 124.3)
	const x = 100_100 // for easy reading
	fmt.Println(a, x)
}
