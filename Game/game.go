package Game

import (
	"fmt"
	"math/rand"
)

type Click int

const (
	LeftClick Click = iota + 1
	MiddleClick
)

type MinePosition struct {
	Point     Point
	IsCleared bool
}

type Game struct {
	Mines     []MinePosition // расположение мин
	Board     *Board         // игровое поле
	Steps     int            // количество ходов
	MinesLeft int
	GameOver  bool
}

func (g *Game) Init() {
	var board Board
	board.Height = 10
	board.Width = 10

	board.Init()
	g.Mines = []MinePosition{}
	g.MinesLeft = 0
	g.Board = &board
	generateRandomMines(g)
	g.MinesLeft = calculateNotFoundMines(g)
}

func generateRandomMines(g *Game) {
	count := (g.Board.Height * g.Board.Width) / 10
	for i := 0; i < count; i++ {
		generateUniqueMines(g)
	}
}

func generateUniqueMines(g *Game) {
	for {
		notFound := false
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

func Process(g *Game, x int, y int, b Click) {
	if b == LeftClick {
		fmt.Println("moving...")
		g.Open(x, y)
		g.Steps++
		return
	}
	if b == MiddleClick {
		fmt.Println("flagging...")
		g.Board.Flag(x, y)
		g.MinesLeft--
		g.Steps++
		return
	}
}

func (g *Game) Open(x int, y int) {
	currentState := g.Board.Cells[x][y]
	if currentState != Unknown { // do nothing if cell is already opened
		return
	}
	failIfMine(g, x, y)
	openNearbyCells(g, x, y)
}

func (g *Game) Flag(x int, y int) {
	g.Board.Flag(x, y)
}

// проверить не попал ли ход на мину
func failIfMine(g *Game, x int, y int) {
	for _, mine := range g.Mines {
		if mine.Point.x == x && mine.Point.y == y {
			g.GameOver = true
		}
	}
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
