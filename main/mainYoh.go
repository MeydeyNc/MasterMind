package main

import (
	"Mastermind"
	"fmt"
	"strconv"
)

func main() {

	colors := [7]string{"bleu", "vert", "orange", "rouge", "blanc", "rose", "gris"}

	var colorsChosen []string

	for {

		fmt.Println("Voici les couleurs à choisir : bleu, vert, orange, rouge, blanc, rose, gris")
		fmt.Println("Veuillez indiquer les couleurs avec leur numéro d'apparition : 1 pour bleu, 2 pour vert ...")
		fmt.Println("Quelle est la première couleur choisi : ")
		input := Mastermind.Input()

		indexColor, _ := strconv.Atoi(input)

		colorsChosen = append(colorsChosen, colors[indexColor-1])

		fmt.Println("Quelle est la deuxième couleur choisi : ")
		input = Mastermind.Input()

		indexColor, _ = strconv.Atoi(input)

		colorsChosen = append(colorsChosen, colors[indexColor-1])

		fmt.Println("Quelle est la troisième couleur choisi : ")
		input = Mastermind.Input()

		indexColor, _ = strconv.Atoi(input)

		colorsChosen = append(colorsChosen, colors[indexColor-1])

		fmt.Println("Quelle est la quatrième couleur choisi : ")
		input = Mastermind.Input()

		indexColor, _ = strconv.Atoi(input)

		colorsChosen = append(colorsChosen, colors[indexColor-1])

		for i := 0; i < 4; i++ {
			fmt.Print(colorsChosen[i] + " ")
		}

		break

	}
}
