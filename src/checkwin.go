package src

import "fmt"

func VerifierVictoire(partie *Partie, symbole string) bool {
	lignes := len(partie.Grille)
	colonnes := len(partie.Grille[0])
	for i := 0; i < lignes; i++ {
		compteur := 0
		for j := 0; j < colonnes; j++ {
			if partie.Grille[i][j] == symbole {
				compteur++
				if compteur == 4 {
					partie.PartieEnCours = false
					fmt.Printf("%s a gagné (alignement horizontal)\n", symbole)
					return true
				}
			} else {
				compteur = 0
			}
		}
	}
	for j := 0; j < colonnes; j++ {
		compteur := 0
		for i := 0; i < lignes; i++ {
			if partie.Grille[i][j] == symbole {
				compteur++
				if compteur == 4 {
					partie.PartieEnCours = false
					fmt.Printf("%s a gagné (alignement vertical)\n", symbole)
					return true
				}
			} else {
				compteur = 0
			}
		}
	}
	return false
}

func VerifierMatchNul(partie *Partie) bool {
	for i := 0; i < len(partie.Grille); i++ {
		for j := 0; j < len(partie.Grille[0]); j++ {
			if partie.Grille[i][j] == "." {
				return false
			}
		}
	}
	partie.PartieEnCours = false
	fmt.Println("Match nul : la grille est pleine")
	return true
}
