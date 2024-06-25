package Game

type Point struct {
	x int
	y int
}

type Board struct {
	Width  int           `json:"width"`
	Height int           `json:"height"`
	Cells  [][]CellState `json:"cells"`
}

func (b *Board) Init() {
	b.Cells = make([][]CellState, b.Height)
	for i := 0; i < b.Height; i++ {
		b.Cells[i] = make([]CellState, b.Width)
		for j := 0; j < b.Width; j++ {
			b.Cells[i][j] = Unknown
		}
	}
}

func (b *Board) Flag(x int, y int) {
	currentState := b.Cells[x][y]
	if currentState == Flagged {
		b.Cells[x][y] = Unknown
		return
	}
	b.Cells[x][y] = Flagged
}
