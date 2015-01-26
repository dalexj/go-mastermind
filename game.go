package main

import (
	"math/rand"
	"strings"
)

const colors = "rgby"

type Game struct {
	Code    string
	Guesses int
}

func NewGame() *Game {
	code := ""
	for i := 0; i < len(colors); i++ {
		code += string(colors[rand.Intn(len(colors))])
	}
	return &Game{
		Code:    code,
		Guesses: 0,
	}
}

func (g Game) IsValidGuess(guess string) bool {
	if len(guess) != len(g.Code) {
		return false
	}
	for i := 0; i < len(guess); i++ {
		if !strings.Contains(colors, string(guess[i])) {
			return false
		}
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
