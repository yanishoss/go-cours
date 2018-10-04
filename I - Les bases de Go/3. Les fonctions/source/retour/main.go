package main

import (
	"errors"
	"fmt"
)

func estMajeur(age uint) bool {
	return age >= 18
}

func estTropAge(age uint) bool {
	return age >= 50
}

func entrerEnBoite(prenom string, age uint) (string, error) {
	if !estMajeur(age) {
		return "", errors.New("Il faut avoir plus de 18 ans pour entrer en boîte, " + prenom + " !")
	}

	if estTropAge(age) {
		return "", errors.New("Les quinquagénaires ne sont pas acceptés, " + prenom + " !")
	}

	return "Bienvenue en boîte, " + prenom, nil
}

func main() {
	msg, err := entrerEnBoite("Sabrina", 56)
	msg2, err2 := entrerEnBoite("Alexia", 25)

	if err != nil {
		fmt.Println(err)
	}

	if err2 != nil {
		fmt.Println(err2)
	}

	fmt.Println(msg)
	fmt.Println(msg2)
}
