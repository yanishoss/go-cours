/*
	Les types de variables:
	- int8: un entier de −127 à +127
	- int16: un entier de -32767 à +32767
	- int32: un entier de -2147483648 à +2147483647
	- int64: un entier de −9223372036854775808 à +9223372036854775807
	- int: un entier de taille variable (peut aller de int8 à int64)

	- uint8: un entier de 0 à +255
	- uint16: un entier de 0 à +65535
	- uint32: un entier de 0 à +4294967295
	- uint64: un entier de 0 à +18446744073709551615
	- uintptr: un entier contenant un pointeur, encodé différemment suivant la plateforme
	- byte: un autre nom pour uint8
	- uint: un entier non-signé de taille variable (peut aller de uint8 à uint64)

	- float32: un réel de -3.4E+38 à +3.4E+38 avec jusqu'à 7 chiffres après la virgule
	- float64: un réel de -1.7E+308 à +1.7E+308 avec jusqu'à 16 chiffres après la virgule

	- complex64: un complexe encodé sur 64 bits
	- complex128: un complexe encodé sur 128 bits

	- bool: un booléan encodé sur 8 bits, peut avoir la valeur 0 (false) ou la valeur 1 (true)

	- rune: un caractère UTF-8 encodé sur 8 bits
	- string: une chaîne de caractères

*/

package main

import (
	"fmt"
)

func main() {
	var integer1 int8 = 2
	var integer2 int16 = -992
	var integer3 int32 = +66345
	var integer4 int64 = -99364447

	var uinteger1 uint8 = 2
	var byte1 byte = 33 // Pareil que uint8
	var uinteger2 uint16 = 992
	var uinteger3 uint32 = 66345
	var uinteger4 uint64 = 99364447

	var decimal1 float32 = 3.9
	var decimal2 float64 = 2990.992

	var complex1 complex64 = complex(10.0, 9.0)
	var complex2 complex128 = complex(10.0, 990.0)

	var boolean bool = true

	var caracter rune = 'A'

	var message string = "Je suis un message!"

	fmt.Println("integer1 int8: ", integer1)
	fmt.Println("integer2 int16: ", integer2)
	fmt.Println("integer3 int32: ", integer3)
	fmt.Println("integer4 int64: ", integer4)

	fmt.Println("uinteger1 uint8: ", uinteger1)
	fmt.Println("byte1 byte: ", byte1)
	fmt.Println("uinteger2 uint16: ", uinteger2)
	fmt.Println("uinteger3 uint32: ", uinteger3)
	fmt.Println("uinteger4 uint64: ", uinteger4)

	fmt.Println("decimal1 float32: ", decimal1)
	fmt.Println("decimal2 float64: ", decimal2)

	fmt.Println("complex1 complex64: ", complex1)
	fmt.Println("complex2 complex128: ", complex2)

	fmt.Println("boolean bool: ", boolean)

	fmt.Println("caracter rune: ", caracter)

	fmt.Println("message string: ", message)
}
