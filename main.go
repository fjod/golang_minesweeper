package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	G "minesweeper/Game"
	"net/http"
	"strconv"
)

var game G.Game

type GameState struct {
	Board     G.Board `json:"board"`
	MinesLeft int     `json:"minesLeft"`
	Steps     int     `json:"steps"`
	GameOver  bool    `json:"gameOver"`
}

func getState() *GameState {
	var g GameState
	g.Board = *game.Board
	g.MinesLeft = game.MinesLeft
	g.Steps = game.Steps
	g.GameOver = game.GameOver
	return &g
}

func initGame(c *gin.Context) {
	game.Init()
	state := getState()
	c.IndentedJSON(http.StatusOK, state)
}

func processStep(c *gin.Context) {
	x, _ := strconv.Atoi(c.Params[0].Value)
	y, _ := strconv.Atoi(c.Params[1].Value)
	b, _ := strconv.Atoi(c.Params[2].Value)
	click := G.Click(b)
	G.Process(&game, x, y, click)
	state := getState()
	c.IndentedJSON(http.StatusOK, state)
}

func createHttpServer() *gin.Engine {
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

	router.GET("/init", initGame)
	router.GET("/step/:x/:y/:b", processStep)
	return router
}

func main() {

	router := createHttpServer()
	err := router.Run("localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
}
