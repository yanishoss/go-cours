package main

func main() {
	var v1 int = 3
	var v2 int = -6
	var v3 string = "Je suis une string!"
	var v4 float32 = 3.9
	var v5 float64 = 66.3

	v1 = 9                    // Autorisé car 9 est une valeur entière de type "int"
	v1 = -9                   // Autorisé car -9 est une valeur entière de type "int"
	v1 = 9.8                  // Pas autorisé car 9.8 est de type "float"
	v1 = "Je suis une string" // Pas autorisé car "Je suis une string" est de type "string"
	v1 = v2                   // Autorisé car "v2" est de type "int"
	v1 = v3                   // Pas autorisé car "v3" est de type "string"

	v3 = "Je suis une string mutée!" // Autorisé car "Je suis une string mutée!" est de type "string"
	v3 = 9                           // Pas autorisé car 9 est de type "int"

	v4 = 5.363 // Autorisé car 5.363 est une valeur décimale de type "float"
	v4 = v5    // Pas autorisé car "v5" est de type "float64"
}
