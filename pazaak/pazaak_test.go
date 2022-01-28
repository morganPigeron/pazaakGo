package pazaak

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewGame(t *testing.T) {
	game := NewGame()
	assert.NotNil(t, game)
}

func TestAddPlayer(t *testing.T) {
	game := getGameWithOnePlayer()
	assert.Equal(t, len(game.players), 1)
}

func TestDefineFirstPlayer(t *testing.T) {
	game := getGameWithTwoPlayers()
	assert.True(t, game.firstPlayerIndex >= 0 && game.firstPlayerIndex < len(game.players))
}

func TestFromInitialStateToPlayingState(t *testing.T) {
	game := getGameWithTwoPlayers()
	game.start()
	assert.Equal(t, game.state, "playing")
}

// Test Helpers

func getGameWithTwoPlayers() *pazaakGame {
	game := NewGame()
	game.addPlayer()
	game.addPlayer()
	game.defineFirstPlayer()
	return game
}

func getGameWithOnePlayer() *pazaakGame {
	game := NewGame()
	game.addPlayer()
	return game
}
