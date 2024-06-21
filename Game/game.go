package Game

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
	board.Height = 10
	board.Width = 10
	board.Init()
	g.Board = &board
	generateRandomMines(g)
	g.Refresh()
}

func generateRandomMines(g *Game) {
	count := (g.Board.Height * g.Board.Width) / 10
	for i := 0; i < count; i++ {
		generateUniqueMines(g)
	}
}

func generateUniqueMines(g *Game) {
	notFound := true
	for notFound {
		notFound = false
		x := rand.Intn(g.Board.Width)
		y := rand.Intn(g.Board.Height)
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

func (g *Game) Open() {
	clearConsole()
	fmt.Println("use arrows to move, f to flag, space to open cell, x to exit")
	mines := calculateNotFoundMines(g)
	fmt.Println("mines:", mines)
	fmt.Println("steps:", g.Steps)
	fmt.Println("waiting for your step..")

	failIfMine(g)
	openNearbyCells(g)

	g.Board.Print()
}

func (g *Game) Flag() {
	clearConsole()
	fmt.Println("use arrows to move, f to flag, space to open cell, x to exit")
	mines := calculateNotFoundMines(g)
	fmt.Println("mines:", mines)
	fmt.Println("steps:", g.Steps)
	fmt.Println("waiting for your step..")

	failIfMine(g)
	openNearbyCells(g)

	g.Board.Print()
}

// проверить не попал ли ход на мину
func failIfMine(g *Game) {
	position := g.Board.current
	for _, mine := range g.Mines {
		if mine.Point.x == position.x && mine.Point.y == position.y {
			fmt.Println("you stepped on a mine")
			os.Exit(1)
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
