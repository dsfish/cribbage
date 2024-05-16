package main

import (
	"fmt"

	"cribbage/card"
	"cribbage/scoring"
)

func main() {
	c := []card.Card{
		card.New("5 ♣"),
		card.New("5 ♦"),
		card.New("5 ♥"),
		card.New("5 ♠"),
		card.New("J ♦"),
	}
	tricks := scoring.GetTricks(c)
	fmt.Printf("%v\n", tricks)
	fmt.Printf("Score: %v\n", scoring.GetScore(tricks))
}
