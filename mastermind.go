package main

import (
	"bufio"
	"fmt"
					"os"
	"strings"
)

func getInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func main() {
	mastermind := NewGame()
	for {
		input := getInput("Enter a 4 color guess: ")

		if !mastermind.IsValidGuess(input) {
			fmt.Println("That was not a valid guess\n")
			continue
		}

		colorsCorrect, positionsCorrect := mastermind.Guess(input)

		fmt.Println("guesses:", mastermind.Guesses)
		fmt.Printf("You got %d correct positions\n\n", positionsCorrect)
		fmt.Printf("You got %d correct colors\n\n", colorsCorrect)

		if positionsCorrect == len(mastermind.Code) {
			fmt.Println("You win, go home.")
			break
		}
	}
}
