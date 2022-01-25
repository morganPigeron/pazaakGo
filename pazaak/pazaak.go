package pazaak

import (
	"math/rand"
	"time"
)

type game struct {
	BoardCards []Card `json:"boardCards"`
}

type player struct {
	Deck  []Card `json:"deck"`
	Cards []Card `json:"cards"`
	Point int8   `json:"point"`
	Win   int8   `json:"win"`
}

type Card struct {
	Value int8   `json:"value"`
	Type  string `json:"type"`
}

var mainCards = []Card{
	{Value: 1, Type: "normal"},
	{Value: 2, Type: "normal"},
	{Value: 3, Type: "normal"},
	{Value: 4, Type: "normal"},
	{Value: 5, Type: "normal"},
	{Value: 6, Type: "normal"},
	{Value: 7, Type: "normal"},
	{Value: 8, Type: "normal"},
	{Value: 9, Type: "normal"},
	{Value: 10, Type: "normal"},
}

var sideCards = []Card{
	{Value: -6, Type: "side"},
	{Value: -5, Type: "side"},
	{Value: -4, Type: "side"},
	{Value: -3, Type: "side"},
	{Value: -2, Type: "side"},
	{Value: -1, Type: "side"},
	{Value: 6, Type: "side"},
	{Value: 5, Type: "side"},
	{Value: 4, Type: "side"},
	{Value: 3, Type: "side"},
	{Value: 2, Type: "side"},
	{Value: 1, Type: "side"},
	{Value: 6, Type: "switchable"},
	{Value: 5, Type: "switchable"},
	{Value: 4, Type: "switchable"},
	{Value: 3, Type: "switchable"},
	{Value: 2, Type: "switchable"},
	{Value: 1, Type: "switchable"},
}

var Players = []player{
	{
		Cards: []Card{},
		Deck:  []Card{},
		Point: 0,
	},
	{
		Cards: []Card{},
		Deck:  []Card{},
		Point: 0,
	},
}

var GameObject = []game{
	{
		BoardCards: []Card{},
	},
	{
		BoardCards: []Card{},
	},
}

func GetRandomMainCard() Card {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	return mainCards[random.Intn(len(mainCards))]
}

func DrawRandomCardsfromDeck(p *player) {

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	for i := 0; i < 4; i++ {
		p.Cards = append(p.Cards, p.Deck[random.Intn(len(p.Deck))])
	}

}

func CreateNewGame(g *[]game) {
	for i := 0; i < len(*g); i++ {
		(*g)[i].BoardCards = []Card{}
	}
}

func AddCardToBoard(g *game) {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	g.BoardCards = append(g.BoardCards, mainCards[random.Intn(len(mainCards))])
}

func PlayerAddCard(g *game, p *player, cardIndex int8) {
	g.BoardCards = append(g.BoardCards, p.Cards[cardIndex])
	p.Cards = append(p.Cards[:cardIndex], p.Cards[cardIndex+1:]...)
}

func EndTurn(g *game, p *player) {
	p.Point = 0
	for i := 0; i < len(g.BoardCards); i++ {
		p.Point += g.BoardCards[i].Value
	}
}

func AssignDeck(p *player, d []Card) {
	p.Deck = append(p.Deck, d...)
}
