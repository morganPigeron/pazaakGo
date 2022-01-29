package pazaak

import "math/rand"

type pazaakGame struct {
	players           []*player
	currentPlayerTurn int
	state             string
}

func (game *pazaakGame) addPlayer() {
	player := makePlayer()
	game.players = append(game.players, player)
}

func (game *pazaakGame) defineFirstPlayer() {
	game.currentPlayerTurn = rand.Intn(len(game.players))
}

func (game *pazaakGame) start() {
	game.state = "playing"
}
