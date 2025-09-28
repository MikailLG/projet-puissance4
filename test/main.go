package main

import (
	"fmt"
	"monprojet/src"
)

func main() {
	partie := src.InitialiserPartie()

	for partie.PartieEnCours {
		src.AfficherGrille(partie.Grille)
		fmt.Printf("C'est au tour de %s (%s)\n", partie.JoueurCommence.Nom, partie.JoueurCommence.Symbole)

		var colonne int
		fmt.Print("Choisissez une colonne (1-7) : ")
		fmt.Scan(&colonne)

		if !src.AjouterPion(&partie, colonne, partie.JoueurCommence.Symbole) {
			continue
		}
		if src.VerifierVictoire(&partie, partie.JoueurCommence.Symbole) {
			src.AfficherGrille(partie.Grille)
			break
		}
		if partie.JoueurCommence.Symbole == "X" {
			partie.JoueurCommence = partie.Joueur2
		} else {
			partie.JoueurCommence = partie.Joueur1
		}
	}
}
