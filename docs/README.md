# Puissance 4 en Go

Ce projet est une implémentation du jeu Puissance 4 en langage Go.  
Le jeu se joue dans la console entre deux joueurs humains.

## Fonctionnalités

- Initialisation de la grille de jeu (6 lignes x 7 colonnes)
- Création de deux joueurs :
  - Joueur 1 → X
  - Joueur 2 → O
- Affichage de la grille dans la console
- Ajout d’un pion : le pion tombe jusqu’à la première case vide dans la colonne choisie
- Vérification automatique de la victoire (alignement horizontal ou vertical)
- Vérification du match nul (grille pleine sans vainqueur)
- Gestion du tour de jeu entre les deux joueurs


- Tous les fichiers dans `src/` utilisent `package src`.
- Les fonctions exportées commencent par une majuscule.

## Lancer le projet

1. Assurez-vous que Go est installé (version 1.22 ou supérieure)
2. Depuis la racine du projet :

```bash
go run main.go
