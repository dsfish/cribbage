package powerset

import "cribbage/card"

func Get(a card.Cards) []card.Cards {
	if len(a) == 0 {
		return []card.Cards{
			{},
		}
	}

	var result []card.Cards
	for _, sub := range Get(a[1:]) {
		result = append(result, sub)
		result = append(result, append(sub, a[0]))
	}
	return result
}
