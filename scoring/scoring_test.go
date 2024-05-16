package scoring

import (
	"testing"

	"cribbage/card"

	"github.com/stretchr/testify/assert"
)

func TestGetTricks(t *testing.T) {
	testCases := []struct {
		input    card.Cards
		expected tricks
	}{
		{
			input: card.Cards{},
			expected: tricks{
				redJacks: card.Cards{},
				fifteens: []card.Cards{},
				sets:     []card.Cards{},
				runs:     []card.Cards{},
			},
		},
		{
			input: card.Cards{
				{
					card.Spades,
					5,
				},
				{
					card.Hearts,
					5,
				},
				{
					card.Clubs,
					5,
				},
				{
					card.Diamonds,
					5,
				},
				{
					card.Diamonds,
					11,
				},
			},
			expected: tricks{
				redJacks: card.Cards{
					{
						card.Diamonds,
						11,
					},
				},
				fifteens: []card.Cards{
					{
						{
							card.Clubs,
							5,
						},
						{
							card.Hearts,
							5,
						},
						{
							card.Spades,
							5,
						},
					},
					{
						{
							card.Diamonds,
							5,
						},
						{
							card.Hearts,
							5,
						},
						{
							card.Spades,
							5,
						},
					},
					{
						{
							card.Diamonds,
							5,
						},
						{
							card.Clubs,
							5,
						},
						{
							card.Spades,
							5,
						},
					},
					{
						{
							card.Diamonds,
							5,
						},
						{
							card.Clubs,
							5,
						},
						{
							card.Hearts,
							5,
						},
					},
					{
						{
							card.Diamonds,
							11,
						},
						{
							card.Spades,
							5,
						},
					},
					{
						{
							card.Diamonds,
							11,
						},
						{
							card.Hearts,
							5,
						},
					},
					{
						{
							card.Diamonds,
							11,
						},
						{
							card.Clubs,
							5,
						},
					},
					{
						{
							card.Diamonds,
							11,
						},
						{
							card.Diamonds,
							5,
						},
					},
				},
				sets: []card.Cards{
					{
						{
							card.Diamonds,
							5,
						},
						{
							card.Clubs,
							5,
						},
						{
							card.Hearts,
							5,
						},
						{
							card.Spades,
							5,
						},
					},
				},
				runs: []card.Cards{},
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.input.String(), func(t *testing.T) {
			assert.Equal(t, testCase.expected, GetTricks(testCase.input))
		})
	}
}

func TestIsFifteen(t *testing.T) {
	testCases := []struct {
		description string
		input       card.Cards
		expected    bool
	}{
		{
			description: "empty",
			expected:    false,
		},
		{
			description: "1 2",
			input: card.Cards{
				{
					card.Spades,
					1,
				},
				{
					card.Hearts,
					2,
				},
			},
			expected: false,
		},
		{
			description: "3 1 2 9",
			input: card.Cards{
				{
					card.Diamonds,
					3,
				},
				{
					card.Spades,
					1,
				},
				{
					card.Hearts,
					2,
				},
				{
					card.Hearts,
					9,
				},
			},
			expected: true,
		},
		{
			description: "5 5 5",
			input: card.Cards{
				{
					card.Diamonds,
					5,
				},
				{
					card.Spades,
					5,
				},
				{
					card.Hearts,
					5,
				},
			},
			expected: true,
		},
		{
			description: "J 5",
			input: card.Cards{
				{
					card.Spades,
					11,
				},
				{
					card.Hearts,
					5,
				},
			},
			expected: true,
		},
		{
			description: "9 5",
			input: card.Cards{
				{
					card.Spades,
					9,
				},
				{
					card.Hearts,
					5,
				},
			},
			expected: false,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			assert.Equal(t, testCase.expected, isFifteen(testCase.input))
		})
	}
}

func TestIsSet(t *testing.T) {
	testCases := []struct {
		description string
		input       card.Cards
		expected    bool
	}{
		{
			description: "empty",
			expected:    false,
		},
		{
			description: "1 2",
			input: card.Cards{
				{
					card.Spades,
					1,
				},
				{
					card.Hearts,
					2,
				},
			},
			expected: false,
		},
		{
			description: "1 1 2",
			input: card.Cards{
				{
					card.Diamonds,
					1,
				},
				{
					card.Spades,
					1,
				},
				{
					card.Hearts,
					2,
				},
			},
			expected: false,
		},
		{
			description: "3 3 3",
			input: card.Cards{
				{
					card.Diamonds,
					3,
				},
				{
					card.Spades,
					3,
				},
				{
					card.Hearts,
					3,
				},
			},
			expected: true,
		},
		{
			description: "J J",
			input: card.Cards{
				{
					card.Diamonds,
					11,
				},
				{
					card.Spades,
					11,
				},
			},
			expected: true,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			assert.Equal(t, testCase.expected, isSet(testCase.input))
		})
	}
}

func TestIsRun(t *testing.T) {
	testCases := []struct {
		description string
		input       card.Cards
		expected    bool
	}{
		{
			description: "empty",
			expected:    false,
		},
		{
			description: "1 2",
			input: card.Cards{
				{
					card.Spades,
					1,
				},
				{
					card.Hearts,
					2,
				},
			},
			expected: true,
		},
		{
			description: "3 1 2",
			input: card.Cards{
				{
					card.Diamonds,
					3,
				},
				{
					card.Spades,
					1,
				},
				{
					card.Hearts,
					2,
				},
			},
			expected: true,
		},
		{
			description: "3 1 4",
			input: card.Cards{
				{
					card.Diamonds,
					3,
				},
				{
					card.Spades,
					1,
				},
				{
					card.Hearts,
					4,
				},
			},
			expected: false,
		},
		{
			description: "10 J Q K",
			input: card.Cards{
				{
					card.Diamonds,
					10,
				},
				{
					card.Spades,
					11,
				},
				{
					card.Hearts,
					12,
				},
				{
					card.Diamonds,
					13,
				},
			},
			expected: true,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			assert.Equal(t, testCase.expected, isRun(testCase.input))
		})
	}
}
