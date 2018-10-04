package main

import (
	"fmt"
)

func addInt(a ...int) int {
	var total int

	for _, n := range a {
		total += n
	}

	return total
}

func addFloat(a float64, b float64) float64 {
	return a + b
}

func main() {
	fmt.Println(addInt(8, 10, 10, 28, 898, 897, 10, 47, 78/3))
	fmt.Println(addFloat(7.5, 9.5))
}
