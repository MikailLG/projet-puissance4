package src

func AjouterPion(partie *Partie, colonne int, symbole string) bool {
	colonne--
	lignes := len(partie.Grille)
	for i := lignes - 1; i >= 0; i-- {
		if partie.Grille[i][colonne] == "." {
			partie.Grille[i][colonne] = symbole
			partie.NombreTours++
			return true
		}
	}
	return false
}
