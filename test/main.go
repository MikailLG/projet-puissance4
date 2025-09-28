package main

import (
	"fmt"
	"monprojet/src"
)

func main() {
	partie := src.InitialiserPartie()
	src.AfficherGrille(partie.Grille)
	fmt.Printf("C'est au tour du %s (%s)\n",
		partie.JoueurCommence.Nom,
		partie.JoueurCommence.Symbole,
	)
}
