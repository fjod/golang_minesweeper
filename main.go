package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, playground")
	var game Game
	game.Init()

	for {
		time.Sleep(100 * time.Millisecond)
		game.ProcessKeyStroke()
	}

}
