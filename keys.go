package main

import (
	"bufio"
	"fmt"
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

func (g *Game) ProcessKeyStroke() {
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
		g.Board.Move(Up)
		g.Refresh()
		return
	}
	if key == KDown {
		g.Board.Move(Down)
		g.Refresh()
		return
	}
	if key == KLeft {
		g.Board.Move(Left)
		g.Refresh()
		return
	}
	if key == KRight {
		g.Board.Move(Right)
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
	case 0x1B:
		return KUp
	case 0x1A:
		return KDown
	case 0x1C:
		return KLeft
	case 0x1D:
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
