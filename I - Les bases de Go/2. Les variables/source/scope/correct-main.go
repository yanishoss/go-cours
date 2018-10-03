package main

import (
	"fmt"
)

var (
	v3 int = 4
)

func main() {
	v1 := 9
	var v4 bool

	fmt.Println(v1) // Fonctionne car "v1" est dans le même "block"
	fmt.Println(v3) // Fonctionne car "v3" est dans un "block" parent (variable globale)

	for i := 0; i < 3; i++ {
		v4 = true

		v2 := fonction()

		fmt.Println(v2) // Fonctionne car "v2" est dans le même "block"
	}

	fmt.Println(v4) // Fonctionne car "v4" est dans le même "block"
}

func fonction() string {
	v2 := "Je suis une string dans une fonction"

	return v2
}
