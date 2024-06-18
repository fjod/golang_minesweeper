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

func ProcessKeyStroke(g *G.Game) {
	char, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println("Error reading key:", err)
		return
	}
	key := convertRuneToKeyStroke(char)
	if key == KExit {
		fmt.Println("Exiting...")
		os.Exit(0)
	}
	if key == KUp {
		g.Board.Move(G.Left)
		g.Refresh()
		return
	}
	if key == KDown {
		g.Board.Move(G.Right)
		g.Refresh()
		return
	}
	if key == KLeft {
		g.Board.Move(G.Up)
		g.Refresh()
		return
	}
	if key == KRight {
		g.Board.Move(G.Down)
		g.Refresh()
		return
	}
	if key == KFlag {
		fmt.Println("flagging...")
		g.Refresh()
		g.Steps++
		return
	}
	if key == KOpen {
		fmt.Println("opening...")
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
