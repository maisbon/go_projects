package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	
	// Récupérer les arguments de la ligne de commande
	Arg := os.Args[1:]

	// Vérifie si le nombre d'arguments n'est pas égal à 1 ou si l'arguments lui-même est vide 
	if len(Arg) != 1 || Arg[0] == "" {
		return
	}

	// Valider les caractères de l'argument
	for _, s := range Arg[0] {
		if s < ' ' || s > '~' {
			return
		}
	}

	// Lire le contenu du fichier "standard.text"
	bytes, err := os.ReadFile("standard.txt")
	// Si une erreur survient lors de la lecture du fichier, l'erreur est affichée et quitte le programme.
	if err != nil {
		os.Exit(1)
	}
	// Diviser le contenu du fichier par les sauts de ligne pour obtenir un tableau de lignes
	ligne := strings.Split(string(bytes), "\n")
    // Vérife si l'argument est "\n" et imprimer un caractère de saut de ligne
	if Arg[0] == "\\n" {
		printNewlineAsc()
		return
	}
    // Initialiser des variables pour le traitement de la chaine d'entrée
	var str []rune
	Newline := false

    // Parcours chaque caractère de la chaîne d'argument
	for i, s := range Arg[0] {
		if Newline {
	// Si un saut de ligne est rencontré, imprimer les caractères accumulés sous forme d'art ASCII
			Newline = false
			printAsc(str, ligne)
			str = []rune{} // Réinitialiser le tampon de caractères
			continue
		}

    // Vérifie le cas spécial : "\" suivi de "n" pour un caractère de saut de ligne
		if s == '\\' && len(Arg[0]) != i+1 {
			if Arg[0][i+1] == 'n' {
				Newline = true
				continue
			}
		}
        // Ajoutez le caractère actuel à la tranche 'str'.
		str = append(str, s)
	}
    // Imprimez l'ensemble final de caractères sous forme d'art ASCII.
	printAsc(str, ligne)
}

// Fonction pour imprimer la représentation en art ASCII des caractères
func printAsc(str []rune, ligne []string) {
	if len(str) != 0 {
		for j := 1; j <= 8; j++ {			
			 for _, s := range str {
			// Calculez le décalage d'indice pour l'art ASCII du caractère actuel
				skip := (s - 32) * 9

		    // Imprimez la ligne correspondante de l'art ASCII pour le caractère actuel
            // Utilisez la valeur 'skip' pour trouver l'indice de départ correct pour l'art du caractère
            // dans la tranche 'ligne' et ajoutez le numéro de ligne actuel pour imprimer la ligne correcte de l'art.
				fmt.Print(ligne[j+int(skip)])
			}
			fmt.Println()
		}
	} else {
		fmt.Println()
	}
}

func printNewlineAsc() {
	fmt.Println()
}
