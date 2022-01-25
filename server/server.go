package server

import (
	"app/pazaak"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func createGame(c *gin.Context) {
	pazaak.CreateNewGame(&pazaak.GameObject)
	c.IndentedJSON(http.StatusOK, pazaak.GameObject)
}

type DeckResponse struct {
	Deck []pazaak.Card `json:"deck"`
}

func assignDeck(c *gin.Context) {
	playerId, err := strconv.Atoi(c.Param("playerId"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "playerId must be an integer"})
		return
	}

	var deck DeckResponse
	if err := c.BindJSON(&deck); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error binding json"})
		return
	}

	pazaak.AssignDeck(&pazaak.Players[playerId], deck.Deck)
	c.IndentedJSON(http.StatusCreated, gin.H{"deck": deck, "playerId": playerId})
}

func drawFromDeck(c *gin.Context) {
	playerId, err := strconv.Atoi(c.Param("playerId"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "playerId must be an integer"})
		return
	}

	pazaak.DrawRandomCardsfromDeck(&pazaak.Players[playerId])
	c.IndentedJSON(http.StatusOK, pazaak.Players[playerId])
}

func addCardToBoard(c *gin.Context) {
	playerId, err := strconv.Atoi(c.Param("playerId"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "playerId must be an integer"})
		return
	}

	pazaak.AddCardToBoard(&pazaak.GameObject[playerId])
	c.IndentedJSON(http.StatusOK, pazaak.GameObject[playerId])
}

func playerPlayCard(c *gin.Context) {
	playerId, err := strconv.Atoi(c.Param("playerId"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "playerId must be an integer"})
		return
	}

	cardId, err := strconv.Atoi(c.Param("cardId"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "cardId must be an integer"})
		return
	}

	pazaak.PlayerAddCard(&pazaak.GameObject[playerId], &pazaak.Players[playerId], int8(cardId))
	c.IndentedJSON(http.StatusOK, pazaak.GameObject[playerId])
}

func endTurn(c *gin.Context) {
	playerId, err := strconv.Atoi(c.Param("playerId"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "playerId must be an integer"})
		return
	}
	pazaak.EndTurn(&pazaak.GameObject[playerId], &pazaak.Players[playerId])
	c.IndentedJSON(http.StatusOK, pazaak.Players[playerId])
}

func Start() {
	router := gin.Default()
	router.GET("/create", createGame)
	router.POST("/assignDeck/:playerId", assignDeck) // FIXME: check if playerId is valid
	router.GET("/drawFromDeck/:playerId", drawFromDeck)
	router.GET("/addCardToBoard/:playerId", addCardToBoard)
	router.GET("/playCard/:playerId/:cardId", playerPlayCard) // FIXME: check if cardId is valid
	router.GET("/endTurn/:playerId", endTurn)
	router.Run("0.0.0.0:8080")
}
