package pazaak

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddDeckToPlayer(t *testing.T) {
	player := makePlayer()
	player.addDeck(testSideDeck)
	assert.Equal(t, len(player.deck), 10)
}

func TestSelectCardsFromDeck(t *testing.T) {
	player := makePlayer()
	player.addDeck(testSideDeck)
	player.selectCardsFromDeck(4)
	assert.Equal(t, len(player.hand), 4)
	assert.Equal(t, len(player.deck), 6)
}

func TestPlayOneCardFromHand(t *testing.T) {
	player := makePlayer()
	player.addDeck(testSideDeck)
	player.selectCardsFromDeck(4)
	player.playOneCardFromHand(0)
	assert.Equal(t, len(player.hand), 3)
	assert.Equal(t, len(player.board), 1)
}

func TestSetStand(t *testing.T) {
	player := makePlayer()
	player.setStand()
	assert.True(t, player.stand)
}

func TestMakePlayer(t *testing.T) {
	player := makePlayer()
	assert.Equal(t, len(player.deck), 0)
	assert.Equal(t, len(player.hand), 0)
	assert.Equal(t, len(player.board), 0)
	assert.Equal(t, player.winCount, 0)
	assert.False(t, player.stand)
	assert.False(t, player.over)
}

func TestSetOver(t *testing.T) {
	player := makePlayer()
	player.setOver()
	assert.True(t, player.over)
}
