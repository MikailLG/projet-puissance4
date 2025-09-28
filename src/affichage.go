package src

import "fmt"

func AfficherGrille(grille [][]string) {
	fmt.Print("  ")
	for i := 0; i < len(grille[0]); i++ {
		fmt.Printf("%d ", i+1)
	}
	fmt.Println()
	for _, ligne := range grille {
		fmt.Print("| ")
		for _, caseValeur := range ligne {
			fmt.Printf("%s ", caseValeur)
		}
		fmt.Println("|")
	}
	fmt.Print("  ")
	for i := 0; i < len(grille[0]); i++ {
		fmt.Print("--")
	}
	fmt.Println()
}
