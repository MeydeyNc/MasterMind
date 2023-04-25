package main

import (
	color "MasterMind/color"
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	color.Comb()

	const (
		codeLength = 4
	)

	rand.Seed(time.Now().UnixNano())

	code := generateCode()

	fmt.Println("start")
	fmt.Println("The code consists of", codeLength, "digits. Can you guess it?")

	scanner := bufio.NewScanner(os.Stdin)

	guesses := 0

	// game loop
	for {
		fmt.Print("guess: ")
		scanner.Scan()
		guess := scanner.Text()

		// check if valid
		if len(guess) != codeLength {
			fmt.Println("Invalid guess. Please enter a", codeLength, "digit code.")
			continue
		}

		// Check if correct
		guesses++
		if guess == code {
			fmt.Println("Congratulations! You guessed the code in", guesses, "guesses.")
			return
		}
		if guesses == 12 {
			fmt.Println("The games end u lose after using", guesses, ".")
		}
		// number of correct digits and correct positions
		correctDigits, correctPositions := checkGuess(guess, code)

		// consol print
		fmt.Println("Incorrect guess. You got", correctDigits, "digits correct and", correctPositions, "in the correct position.")
	}
}

func generateCode() string {
	code := ""
	for i := 0; i < codeLength; i++ {
		digit := rand.Intn(10)
		code += fmt.Sprint(digit)
	}
	return code
}

func checkGuess(guess string, code string) (int, int) {
	correctDigits := 0
	correctPositions := 0
	for i, guessDigit := range guess {
		for j, codeDigit := range code {
			if guessDigit == codeDigit {
				correctDigits++
				if i == j {
					correctPositions++
				}
			}
		}
	}
	return correctDigits, correctPositions
}
