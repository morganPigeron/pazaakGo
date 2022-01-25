package pazaak

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewGame(t *testing.T) {
	CreateNewGame(&GameObject)
	assert.Equal(t, len(GameObject[0].BoardCards), 0)
	assert.Equal(t, len(GameObject[1].BoardCards), 0)

}

func TestDeckAssignation(t *testing.T) {
	newDeck := []Card{
		sideCards[0],
		sideCards[1],
		sideCards[2],
		sideCards[3],
		sideCards[4],
		sideCards[5],
		sideCards[6],
		sideCards[7],
		sideCards[8],
		sideCards[9],
	}

	AssignDeck(&Players[0], newDeck)

	assert.Equal(t, newDeck, Players[0].Deck)
}

func TestHandCardDefinition(t *testing.T) {
	newDeck := []Card{
		sideCards[0],
		sideCards[1],
		sideCards[2],
		sideCards[3],
		sideCards[4],
		sideCards[5],
		sideCards[6],
		sideCards[7],
		sideCards[8],
		sideCards[9],
	}
	Players[0].Deck = newDeck

	DrawRandomCardsfromDeck(&Players[0])

	assert.Equal(t, len(Players[0].Cards), 4)
}

func TestTurnAddCard(t *testing.T) {
	AddCardToBoard(&GameObject[0])
	assert.Equal(t, len(GameObject[0].BoardCards), 1)
}

func TestPlayerAddCard(t *testing.T) {
	hand := []Card{
		sideCards[0],
		sideCards[1],
		sideCards[2],
		sideCards[3],
	}

	want := []Card{
		sideCards[0],
		sideCards[2],
		sideCards[3],
	}

	Players[0].Cards = hand
	GameObject[0].BoardCards = []Card{}
	PlayerAddCard(&GameObject[0], &Players[0], 1)
	assert.Equal(t, Players[0].Cards, want)
	assert.Equal(t, GameObject[0].BoardCards, []Card{sideCards[1]})
}

func TestEndTurn(t *testing.T) {
	GameObject[0].BoardCards = []Card{
		mainCards[0],
		mainCards[0],
		mainCards[4],
	}

	EndTurn(&GameObject[0], &Players[0])
	assert.Equal(t, Players[0].Point, int8(7))
}
