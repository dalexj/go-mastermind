package main

const colors = "rgby"

type game struct {
	code    string
	guesses int
}

func (g game) isValidGuess(guess string) bool {
	if len(guess) != len(g.code) {
		return false
	}
	for i := 0; i < len(guess); i++ {
		// if colors //if !colors.include? guess[i] return false
	}
	return true
}

func (g game) guess(guess string) int {
	g.guesses++
	return g.numCorrect(guess)
}

func (g game) numCorrect(guess string) int {
	if len(g.code) != len(guess) {
		return 0
	}
	count := 0
	for i := 0; i < len(g.code); i++ {
		if g.code[i] == guess[i] {
			count++
		}
	}
	return count
}
