package golang_minesweeper

type Point struct {
	x int
	y int
}

type Board struct {
	width   int
	height  int
	cells   [][]CellState
	current Point
}

type Direction int

const (
	None Direction = iota
	Up
	Down
	Left
	Right
)

func (b *Board) Move(direction Direction) {
	switch direction {
	case None:
		return
	case Up:
		if b.current.y > 0 {
			b.current.y = b.current.y - 1
		}
	case Down:
		if b.current.y < b.height-1 {
			b.current.y = b.current.y + 1
		}
	case Left:
		if b.current.x > 0 {
			b.current.x = b.current.x - 1
		}
	case Right:
		if b.current.x < b.width-1 {
			b.current.x = b.current.x + 1
		}
	default:
		panic("unhandled case in board move")
	}
}
