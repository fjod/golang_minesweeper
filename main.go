package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	G "minesweeper/Game"
	"net/http"
	"strconv"
)

var game G.Game

func getBoard(c *gin.Context) {
	gameState := struct {
		Board     G.Board `json:"board"`
		MinesLeft int     `json:"minesLeft"`
		Steps     int     `json:"steps"`
		GameOver  bool    `json:"gameOver"`
	}{
		Board:     *game.Board,
		MinesLeft: game.MinesLeft,
		Steps:     game.Steps,
		GameOver:  game.GameOver,
	}
	c.IndentedJSON(http.StatusOK, gameState)
}

func processStep(c *gin.Context) {
	x, _ := strconv.Atoi(c.Params[0].Value)
	y, _ := strconv.Atoi(c.Params[1].Value)
	b, _ := strconv.Atoi(c.Params[2].Value)
	G.Process(&game, x, y, b)
	gameState := struct {
		Board     G.Board `json:"board"`
		MinesLeft int     `json:"minesLeft"`
		Steps     int     `json:"steps"`
		GameOver  bool    `json:"gameOver"`
	}{
		Board:     *game.Board,
		MinesLeft: game.MinesLeft,
		Steps:     game.Steps,
		GameOver:  game.GameOver,
	}
	c.IndentedJSON(http.StatusOK, gameState)
}
func main() {

	game.Init()

	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	router.GET("/init", getBoard)
	router.GET("/step/:x/:y/:b", processStep)

	router.Run("localhost:8080")

	fmt.Println("Hello, playground")
}
