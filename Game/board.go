package Game

import "fmt"

type Point struct {
	x int
	y int
}

type Board struct {
	Width     int           `json:"width"`
	Height    int           `json:"height"`
	Cells     [][]CellState `json:"cells"`
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
	b.Cells[b.current.x][b.current.y] = b.prevState   // restore prev state to current cell
	f(b)                                              // move
	b.prevState = b.Cells[b.current.x][b.current.y]   // save current state for next move
	b.Cells[b.current.x][b.current.y] = Selected_cell // select current cell
}

func (b *Board) Move2(x int, y int) {
	if x < 0 || x >= b.Width || y < 0 || y >= b.Height {
		return
	}
	b.current.x = x
	b.current.y = y
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
		if b.current.y < b.Height-1 {
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
		if b.current.x < b.Width-1 {
			localMove(b, func(b *Board) {
				b.current.x = b.current.x + 1
			})
		}
	default:
		panic("unhandled case in board move")
	}
}

func (b *Board) Print() {
	for i := 0; i < b.Height; i++ {
		for j := 0; j < b.Width; j++ {
			fmt.Printf("%c", b.Cells[i][j].GetRune())
		}
		fmt.Println() // new line
	}
}

func (b *Board) Init() {
	b.Cells = make([][]CellState, b.Height)
	for i := 0; i < b.Height; i++ {
		b.Cells[i] = make([]CellState, b.Width)
		for j := 0; j < b.Width; j++ {
			b.Cells[i][j] = Unknown
		}
	}
	b.prevState = Unknown
	b.Cells[0][0] = Selected_cell
}

func (b *Board) Flag(x int, y int) {
	b.Cells[x][y] = Flagged
}
