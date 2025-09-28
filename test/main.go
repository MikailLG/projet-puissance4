package main

import (
	"fmt"
	"monprojet/src"
)

func main() {
	partie := src.InitialiserPartie()
	src.AfficherGrille(partie.Grille)
	fmt.Println("\n👉 Joueur 1 joue dans la colonne 4")
	src.AjouterPion(&partie, 4, partie.Joueur1.Symbole)
	src.AfficherGrille(partie.Grille)
	fmt.Println("\n👉 Joueur 2 joue dans la colonne 4")
	src.AjouterPion(&partie, 4, partie.Joueur2.Symbole)
	src.AfficherGrille(partie.Grille)
	fmt.Println("\n👉 Joueur 1 joue dans la colonne 1")
	src.AjouterPion(&partie, 1, partie.Joueur1.Symbole)
	src.AfficherGrille(partie.Grille)
}
