package src

type Joueur struct {
	Nom     string
	Symbole string
}

type Partie struct {
	Grille          [][]string
	Joueur1         Joueur
	Joueur2         Joueur
	JoueurCommence  Joueur
	NombreTours     int
	PartieEnCours   bool
}

func InitialiserPartie() Partie {
	lignes := 6
	colonnes := 7
	grille := make([][]string, lignes)

	for i := 0; i < lignes; i++ {
		grille[i] = make([]string, colonnes)
		for j := 0; j < colonnes; j++ {
			grille[i][j] = "."
		}
	}

	joueur1 := Joueur{Nom: "Joueur 1", Symbole: "X"}
	joueur2 := Joueur{Nom: "Joueur 2", Symbole: "O"}

	return Partie{
		Grille:         grille,
		Joueur1:        joueur1,
		Joueur2:        joueur2,
		JoueurCommence: joueur1,
		NombreTours:    0,
		PartieEnCours:  true,
	}
}
