package deck

import "cribbage/card"

type Deck struct {
	cards card.Cards
}

func New() Deck {
	cards := make(card.Cards, 52)
	i := 0
	for _, suit := range []card.Suit{
		card.Spades,
		card.Clubs,
		card.Hearts,
		card.Diamonds,
	} {
		for value := 1; value < 14; value++ {
			cards[i] = card.Card{
				Suit:  suit,
				Value: value,
			}
			i++
		}
	}
	return Deck{cards}
}
