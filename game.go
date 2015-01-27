package main

import (
	"math/rand"
	"strings"
	"time"
)

const colors = "rgby"

type Game struct {
	Code    string
	Guesses int
}

func NewGame() *Game {
	code := ""
	for i := 0; i < len(colors); i++ {
		code += string(colors[randNum(len(colors))])
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

func (g *Game) Guess(guess string) (int, int) {
	g.Guesses++
	return g.ColorsCorrect(guess), g.PositionsCorrect(guess)
}

func (g Game) PositionsCorrect(guess string) int {
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

func (g Game) ColorsCorrect(guess string) int {
	count := 0
	for i, _ := range g.Code {
		char := string(g.Code[i])
		if strings.Contains(guess, char) {
			count++
			guess = strings.Replace(guess, char, "", 1)
		}
	}
	return count
}

func randNum(max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max)
}
