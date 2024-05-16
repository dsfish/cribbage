package card

import (
	"fmt"
	"strconv"
	"strings"
)

type Color string
type Suit string

const (
	Red   Color = "red"
	Black Color = "black"

	Spades   Suit = "♠"
	Clubs    Suit = "♣"
	Diamonds Suit = "♦"
	Hearts   Suit = "♥"
)

type Card struct {
	Suit  Suit
	Value int
}

func New(s string) Card {
	parts := strings.Split(s, " ")
	v := 0
	switch vRaw := parts[0]; vRaw {
	case "J":
		v = 11
	case "Q":
		v = 12
	case "K":
		v = 13
	case "A":
		v = 1
	default:
		var err error
		v, err = strconv.Atoi(vRaw)
		if err != nil {
			panic(err)
		}
	}

	return Card{
		Suit:  Suit(parts[1]),
		Value: v,
	}
}

func (c Card) String() string {
	var displayValue string
	switch c.Value {
	case 1:
		displayValue = "A"
	case 11:
		displayValue = "J"
	case 12:
		displayValue = "Q"
	case 13:
		displayValue = "K"
	default:
		displayValue = fmt.Sprint(c.Value)
	}
	return displayValue + fmt.Sprint(c.Suit)
}

func (c Card) Color() Color {
	if c.Suit == Spades || c.Suit == Clubs {
		return Black
	}
	return Red
}

type Cards []Card

func Intersects(a, b Cards) bool {
	for _, c1 := range a {
		for _, c2 := range b {
			if c1 == c2 {
				return true
			}
		}
	}
	return false
}

func (c Cards) String() string {
	if len(c) == 0 {
		return ""
	}
	s := "["
	for i := 0; i < len(c)-1; i++ {
		s += c[i].String() + ", "
	}
	s += c[len(c)-1].String() + "]"
	return s
}

type SliceOfCards []Cards

func (s SliceOfCards) String() string {
	cardStrings := make([]string, len(s))
	for i, cards := range s {
		cardStrings[i] = cards.String()
	}
	return strings.Join(cardStrings, ", ")
}
