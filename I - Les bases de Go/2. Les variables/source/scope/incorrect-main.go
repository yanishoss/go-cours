package main

import (
	"fmt"
)

var (
	v3 int = 4
)

func main() {
	v1 := 9

	fmt.Println(v1) // Fonctionne car "v1" est dans le même "block"
	fmt.Println(v3) // Fonctionne car "v3" est dans un "block" parent (variable globale)

	for i := 0; i < 3; i++ {
		v4 := true

		fonction()

		fmt.Println(v2) // Ne fonctionne pas car "v2" n'est pas dans le même "block"
	}

	fmt.Println(v4) // Ne fonctionne pas car "v4" est pas dans le même "block"
}

func fonction() {
	v2 := "Je suis une string dans une fonction"
}
