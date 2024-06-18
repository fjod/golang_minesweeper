package main

import (
	"time"
)

func main() {
	var game Game
	game.Init()

	for {
		time.Sleep(100 * time.Millisecond)
		game.ProcessKeyStroke()
	}
}
