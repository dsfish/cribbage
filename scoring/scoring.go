package scoring

import (
	"fmt"
	"sort"
	"strings"

	"cribbage/card"
	"cribbage/powerset"
)

type tricks struct {
	redJacks card.Cards
	fifteens []card.Cards
	sets     []card.Cards
	runs     []card.Cards
}

func (t tricks) String() string {
	return strings.Join([]string{
		fmt.Sprintf("Red Jacks: %v", t.redJacks),
		fmt.Sprintf("Fifteens: %v", card.SliceOfCards(t.fifteens)),
		fmt.Sprintf("Sets: %v", card.SliceOfCards(t.sets)),
		fmt.Sprintf("Runs: %v", card.SliceOfCards(t.runs)),
	}, "\n")
}

func GetTricks(hand card.Cards) tricks {
	tricks := tricks{
		redJacks: card.Cards{},
		fifteens: []card.Cards{},
		sets:     []card.Cards{},
		runs:     []card.Cards{},
	}

	combinations := powerset.Get(hand)

	// Reverse sort (descending) to get larger sets/runs first
	sort.Slice(combinations, func(i, j int) bool {
		return len(combinations[i]) > len(combinations[j])
	})

	cardsAlreadyInSets := card.Cards{}
	cardsAlreadyInRuns := card.Cards{}

	for i := range combinations {
		c := append(card.Cards{}, combinations[i]...)

		switch len(c) {
		case 0:
			continue
		case 1:
			if isRedJack(c[0]) {
				tricks.redJacks = append(tricks.redJacks, c[0])
			}
		default:
			if isFifteen(c) {
				tricks.fifteens = append(tricks.fifteens, c)
			}
			// Proceed if no overlap with cards already in sets/runs
			if !card.Intersects(c, cardsAlreadyInSets) && isSet(c) {
				tricks.sets = append(tricks.sets, c)
				cardsAlreadyInSets = append(cardsAlreadyInSets, c...)
			}
			if !card.Intersects(c, cardsAlreadyInRuns) && isRun(c) {
				tricks.runs = append(tricks.runs, c)
				cardsAlreadyInRuns = append(cardsAlreadyInRuns, c...)
			}
		}
	}

	return tricks
}

func isRedJack(c card.Card) bool {
	return c.Value == 11 && (c.Suit == card.Diamonds || c.Suit == card.Hearts)
}

func isFifteen(cards card.Cards) bool {
	sum := 0
	for _, c := range cards {
		sum += min(c.Value, 10)
	}
	return sum == 15
}

func isSet(cards card.Cards) bool {
	if len(cards) < 2 {
		return false
	}

	m := map[int]bool{}
	for _, c := range cards {
		m[c.Value] = true
	}
	return len(m) == 1
}

func isRun(cards card.Cards) bool {
	if len(cards) < 3 {
		return false
	}

	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Value < cards[j].Value
	})

	lastCard := cards[0]
	for i := 1; i < len(cards); i++ {
		if cards[i].Value-lastCard.Value != 1 {
			return false
		}
		lastCard = cards[i]
	}
	return true
}

func GetScore(t tricks) int {
	score := 0
	score += len(t.redJacks)
	score += len(t.fifteens) * 2

	for _, run := range t.runs {
		score += len(run)
	}
	for _, set := range t.sets {
		switch len(set) {
		case 2:
			score += 2
		case 3:
			score += 6
		case 4:
			score += 12
		}
	}
	return score
}
