package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	G "minesweeper/Game"
	K "minesweeper/Input"
	"net/http"
	"strconv"
)

var game G.Game

func getBoard(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, &game.Board)
}

func processStep(c *gin.Context) {
	x, _ := strconv.Atoi(c.Query("x"))
	y, _ := strconv.Atoi(c.Query("y"))
	b, _ := strconv.Atoi(c.Query("b"))
	K.ProcessKeyStroke(&game, x, y, b)
	c.IndentedJSON(http.StatusOK, &game.Board)
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
