package main

import (
	"fmt"
	G "minesweeper/Game"
	K "minesweeper/Input"
	"time"
)

func main() {
	fmt.Println("Hello, playground")
	var game G.Game
	game.Init()
	for {
		time.Sleep(100 * time.Millisecond)
		K.ProcessKeyStroke(&game)
	}

}
