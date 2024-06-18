package Game

import "fmt"

type Point struct {
	x int
	y int
}

type Board struct {
	width     int
	height    int
	cells     [][]CellState
	current   Point
	prevState CellState
}

type Direction int

const (
	None Direction = iota
	Up
	Down
	Left
	Right
)

type movementFunc func(b *Board)

func localMove(b *Board, f movementFunc) {
	b.cells[b.current.x][b.current.y] = b.prevState   // restore prev state to current cell
	f(b)                                              // move
	b.prevState = b.cells[b.current.x][b.current.y]   // save current state for next move
	b.cells[b.current.x][b.current.y] = Selected_cell // select current cell
}

func (b *Board) Move(direction Direction) {
	switch direction {
	case None:
		return
	case Up:
		if b.current.y > 0 {
			localMove(b, func(b *Board) {
				b.current.y = b.current.y - 1
			})
		}
	case Down:
		if b.current.y < b.height-1 {
			localMove(b, func(b *Board) {
				b.current.y = b.current.y + 1
			})
		}
	case Left:
		if b.current.x > 0 {
			localMove(b, func(b *Board) {
				b.current.x = b.current.x - 1
			})

		}
	case Right:
		if b.current.x < b.width-1 {
			localMove(b, func(b *Board) {
				b.current.x = b.current.x + 1
			})
		}
	default:
		panic("unhandled case in board move")
	}
}

func (b *Board) Print() {
	for i := 0; i < b.height; i++ {
		for j := 0; j < b.width; j++ {
			fmt.Printf("%c", b.cells[i][j].GetRune())
		}
		fmt.Println() // new line
	}
}

func (b *Board) Init() {
	b.cells = make([][]CellState, b.height)
	for i := 0; i < b.height; i++ {
		b.cells[i] = make([]CellState, b.width)
		for j := 0; j < b.width; j++ {
			b.cells[i][j] = Unknown
		}
	}
	b.prevState = Unknown
	b.cells[0][0] = Selected_cell
}
