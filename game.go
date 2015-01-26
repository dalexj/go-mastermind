package main

import (
	"math/rand"
)

type Game struct {
	Code    string
	Guesses int
}

func NewGame() *Game {
	colors := []string{"r", "g", "b", "y"}
	code := ""
	for i := 0; i < len(colors); i++ {
		code += colors[rand.Intn(len(colors))]
	}
	return &Game{
		Code: code,
		Guesses: 0,
	}
}

func (g Game) IsValidGuess(guess string) bool {
	if len(guess) != len(g.Code) {
		return false
	}
	for i := 0; i < len(guess); i++ {
		// if colors //if !colors.include? guess[i] return false
	}
	return true
}

func (g *Game) Guess(guess string) int {
	g.Guesses++
	return g.NumCorrect(guess)
}

func (g Game) NumCorrect(guess string) int {
	if len(g.Code) != len(guess) {
		return 0
	}
	count := 0
	for i, _ := range g.Code {
		if g.Code[i] == guess[i] {
			count++
		}
	}
	return count
}
