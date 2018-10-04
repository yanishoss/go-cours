package main

import (
	"fmt"
)

func saluer(prenom string) string {
	return "Bonjour, " + prenom + " !"
}

func accueil(prenom string, lieu string) string {
	return "Bonjour, " + prenom + " !\n" + "Bienvenue dans " + lieu + "!"
}

// Exercice:
// Créer une fonction qui dis au-revoir à une personne

func main() {
	fmt.Println(saluer("Jean-Pierre"))
	fmt.Println(accueil("Uxio", "Cathédrale"))
}
