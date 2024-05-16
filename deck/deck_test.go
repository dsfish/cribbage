package deck

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDeck(t *testing.T) {
	expected := []string{
		"A♠",
		"2♠",
		"3♠",
		"4♠",
		"5♠",
		"6♠",
		"7♠",
		"8♠",
		"9♠",
		"10♠",
		"J♠",
		"Q♠",
		"K♠",
		"A♣",
		"2♣",
		"3♣",
		"4♣",
		"5♣",
		"6♣",
		"7♣",
		"8♣",
		"9♣",
		"10♣",
		"J♣",
		"Q♣",
		"K♣",
		"A♥",
		"2♥",
		"3♥",
		"4♥",
		"5♥",
		"6♥",
		"7♥",
		"8♥",
		"9♥",
		"10♥",
		"J♥",
		"Q♥",
		"K♥",
		"A♦",
		"2♦",
		"3♦",
		"4♦",
		"5♦",
		"6♦",
		"7♦",
		"8♦",
		"9♦",
		"10♦",
		"J♦",
		"Q♦",
		"K♦",
	}
	deck := New()
	actual := make([]string, len(deck.cards))
	for i, card := range deck.cards {
		actual[i] = card.String()
	}
	assert.ElementsMatch(t, expected, actual)
}
