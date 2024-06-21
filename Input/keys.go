package Input

import (
	"bufio"
	"fmt"
	G "minesweeper/Game"
	"os"
)

type KeyStroke int

var reader = bufio.NewReader(os.Stdin)

const (
	Other KeyStroke = iota
	KUp
	KDown
	KLeft
	KRight
	KFlag
	KOpen
	KExit
)

func ProcessKeyStroke(g *G.Game, x int, y int, b int) {

	if b == 1 {
		fmt.Println("moving...")
		g.Board.Move2(x, y)
		g.Open()
		g.Refresh()
		g.Steps++
		return
	}
	if b == 2 {
		fmt.Println("flagging...")
		g.Board.Flag(x, y)
		g.Refresh()
		g.Steps++
		return
	}
}

func convertRuneToKeyStroke(char rune) KeyStroke {
	switch char {
	case 119:
		return KUp
	case 115:
		return KDown
	case 97:
		return KLeft
	case 100:
		return KRight
	case 0x66:
		return KFlag
	case 0x20:
		return KOpen
	case 0x78:
		return KExit
	default:
		return Other
	}
}
