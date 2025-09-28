package main

import (
	"fmt"
	"monprojet/src"
)

func main() {
	partie := src.InitialiserPartie()

	for partie.PartieEnCours {
		src.AfficherGrille(partie.Grille)
		var colonne int

		fmt.Printf("\n%s (%s), choisissez une colonne (1-7) : ", partie.JoueurCommence.Nom, partie.JoueurCommence.Symbole)
		fmt.Scan(&colonne)

		if colonne < 1 || colonne > 7 {
			fmt.Println("Colonne invalide. Réessayez.")
			continue
		}
		if !src.AjouterPion(&partie, colonne, partie.JoueurCommence.Symbole) {
			fmt.Println("Colonne pleine. Réessayez.")
			continue
		}
		if src.VerifierVictoire(&partie, partie.JoueurCommence.Symbole) {
			src.AfficherGrille(partie.Grille)
			fmt.Printf("Partie terminée : %s gagne !\n", partie.JoueurCommence.Nom)
			break
		}
		if src.VerifierMatchNul(&partie) {
			src.AfficherGrille(partie.Grille)
			fmt.Println("Partie terminée : match nul.")
			break
		}
		if partie.JoueurCommence.Symbole == partie.Joueur1.Symbole {
			partie.JoueurCommence = partie.Joueur2
		} else {
			partie.JoueurCommence = partie.Joueur1
		}
	}
}
