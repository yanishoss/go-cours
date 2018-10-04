package main

import (
	"fmt"
)

const (
	Red = iota
	Blue
	Green
	Yellow
)

func main() {
	const str = "Salut !"

	fmt.Println(str) // Fonctionne car une constante peut-être lue

	str = "Salut les amis !" // Ne fonctionne pas car une constante ne peut pas être écrite

}
