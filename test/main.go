package main

import (
	"fmt"
	"monprojet/src"
)

func main() {
	partie := src.InitialiserPartie()
	for _, ligne := range partie.Grille {
		fmt.Println(ligne)
	}
	fmt.Printf("C'est au tour de %s (%s)\n",
		partie.JoueurCommence.Nom,
		partie.JoueurCommence.Symbole,
	)
}
