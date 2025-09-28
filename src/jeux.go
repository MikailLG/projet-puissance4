package src

import "fmt"

func AjouterPion(partie *Partie, colonne int, symbole string) bool {
	colonne--

	if colonne < 0 || colonne >= len(partie.Grille[0]) {
		fmt.Println("Colonne invalide. Choisissez un numéro entre 1 et 7.")
		return false
	}
	if partie.Grille[0][colonne] != "." {
		fmt.Println("Cette colonne est déjà pleine. Choisissez une autre.")
		return false
	}
	for i := len(partie.Grille) - 1; i >= 0; i-- {
		if partie.Grille[i][colonne] == "." {
			partie.Grille[i][colonne] = symbole
			return true
		}
	}
	return false
}
