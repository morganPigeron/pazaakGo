package pazaak

import "math/rand"

type player struct {
	deck     []card
	hand     []card
	board    []card
	winCount int
	stand    bool
	over     bool
}

func (p *player) addDeck(deck []card) {
	p.deck = append(p.deck, deck...)
}

func (p *player) selectCardsFromDeck(number int) {
	for i := 0; i < number; i++ {
		card := p.deck[rand.Intn(len(p.deck))]
		p.deck = removeCardFromDeck(p.deck, card)
		p.hand = append(p.hand, card)
	}
}

func (p *player) playOneCardFromHand(index int) {
	card := p.hand[index]
	p.hand = append(p.hand[:index], p.hand[index+1:]...)
	p.board = append(p.board, card)
}

func (p *player) setStand() {
	p.stand = true
}

func (p *player) setOver() {
	p.over = true
}
