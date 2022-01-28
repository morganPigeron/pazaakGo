package pazaak

func NewGame() *pazaakGame {
	return &pazaakGame{
		state: "initial",
	}
}

func makePlayer() *player {
	return &player{
		deck:     []card{},
		hand:     []card{},
		board:    []card{},
		winCount: 0,
		stand:    false,
		over:     false,
	}
}

func removeCardFromDeck(deck []card, card card) []card {
	for i, c := range deck {
		if c.value == card.value {
			deck = append(deck[:i], deck[i+1:]...)
		}
	}
	return deck
}
