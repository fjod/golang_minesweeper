package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
)

type MinePosition struct {
	Point     Point
	IsCleared bool
}

type Game struct {
	Mines []MinePosition // расположение мин
	Board *Board         // игровое поле
	Steps int            // количество ходов
}

func (g *Game) Init() {
	var board Board
	board.height = 10
	board.width = 10
	board.Init()
	g.Board = &board
	generateRandomMines(g)
}

func generateRandomMines(g *Game) {
	count := (g.Board.height * g.Board.width) / 10
	for i := 0; i < count; i++ {
		generateUniqueMines(g)
	}
}

func generateUniqueMines(g *Game) {
	notFound := true
	for notFound {
		notFound = false
		x := rand.Intn(g.Board.width)
		y := rand.Intn(g.Board.height)
		for _, mine := range g.Mines {
			if mine.Point.x == x && mine.Point.y == y {
				notFound = true
				break
			}
		}
		if !notFound {
			g.Mines = append(g.Mines, MinePosition{Point{x, y}, false})
			return
		}
	}
}

func (g *Game) Refresh() {
	clearConsole()
	fmt.Println("use arrows to move, f to flag, space to open cell, x to exit")
	mines := calculateNotFoundMines(g)
	fmt.Println("mines:", mines)
	fmt.Println("steps:", g.Steps)
	fmt.Println("waiting for your step..")
	g.Board.Print()
}

func calculateNotFoundMines(g *Game) int {
	notFoundMines := 0
	for _, mine := range g.Mines {
		if mine.IsCleared == false {
			notFoundMines += 1
		}
	}
	return notFoundMines
}

func clearConsole() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Println("cant clear console")
	}
}
