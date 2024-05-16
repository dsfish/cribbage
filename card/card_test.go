package card

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	testCases := []struct {
		value  int
		suit   Suit
		output string
	}{
		{
			1,
			Spades,
			"A♠",
		},
		{
			2,
			Clubs,
			"2♣",
		},
		{
			3,
			Diamonds,
			"3♦",
		},
		{
			13,
			Hearts,
			"K♥",
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.output, func(t *testing.T) {
			card := Card{
				Value: testCase.value,
				Suit:  testCase.suit,
			}
			assert.Equal(t, testCase.output, card.String())
		})
	}
}

func TestIntersects(t *testing.T) {
	testCases := []struct {
		description string
		a           Cards
		b           Cards
		expected    bool
	}{
		{
			description: "both empty",
			expected:    false,
		},
		{
			description: "no intersection",
			a: []Card{
				{
					Spades,
					1,
				},
				{
					Hearts,
					2,
				},
			},
			b: []Card{
				{
					Spades,
					3,
				},
				{
					Hearts,
					4,
				},
			},
			expected: false,
		},
		{
			description: "intersection",
			a: []Card{
				{
					Spades,
					1,
				},
				{
					Hearts,
					2,
				},
			},
			b: []Card{
				{
					Spades,
					3,
				},
				{
					Hearts,
					2,
				},
			},
			expected: true,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			assert.Equal(t, testCase.expected, Intersects(testCase.a, testCase.b))
		})
	}
}
