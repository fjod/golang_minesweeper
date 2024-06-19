package Game

type CellCheckedState struct {
	checked bool
	x       int
	y       int
	state   CellState
}

func openNearbyCells(g *Game) {
	// проверить 9 соседних клеток и выставить кол-во мин, которых касается каждая клетка
	// если в соседней клетке 0 мин, то поставить её в открытую
	// рекурсивно продолжать пока не закончатся непроверенные клетки

	// сначала создать массив клеток с начальным состоянием
	h := g.Board.height
	w := g.Board.width
	state := make([][]CellCheckedState, h*w)
	for i := 0; i < h; i++ {
		state[i] = make([]CellCheckedState, w)
		for j := 0; j < w; j++ {
			state[i][j] = CellCheckedState{
				state:   g.Board.cells[i][j],
				checked: false,
				x:       i,
				y:       j,
			}
		}
	}

	// без рекурсии найти соседние клетки
	surround := getNearbyCells(&state, g.Board.current.x, g.Board.current.y)
	for _, cell := range surround {
		cell.checked = true
		localMines := getNearbyMinesCount(g, cell.x, cell.y)
		if localMines > 0 {
			g.Board.cells[cell.x][cell.y] = FromCount(localMines) // установить кол-во мин в клетке
		}
		if localMines == 0 {
			// рекурсивно продолжать пока не закончатся непроверенные клетки
		}
	}
}

func getNearbyCells(s *[][]CellCheckedState, x int, y int) []CellCheckedState {
	ret := make([]CellCheckedState, 0)
	w := len(*s)
	h := len((*s)[0])
	appendCellIfPossible(&ret, s, x-1, y-1, w, h)
	appendCellIfPossible(&ret, s, x, y-1, w, h)
	appendCellIfPossible(&ret, s, x+1, y-1, w, h)
	appendCellIfPossible(&ret, s, x-1, y+1, w, h)
	appendCellIfPossible(&ret, s, x, y+1, w, h)
	appendCellIfPossible(&ret, s, x+1, y+1, w, h)
	appendCellIfPossible(&ret, s, x+1, y, w, h)
	appendCellIfPossible(&ret, s, x-1, y, w, h)
	return ret
}

func appendCellIfPossible(ret *[]CellCheckedState, s *[][]CellCheckedState, x int, y int, w int, h int) {
	if x >= 0 && x < w && y >= 0 && y < h {
		cell := (*s)[x][y]
		if cell.checked == false {
			*ret = append(*ret, (*s)[x][y])
		}
	}
}

func getNearbyMinesCount(g *Game, x int, y int) int {
	count := 0
	for _, c := range g.Mines {
		if c.Point.x == x-1 && c.Point.y == y-1 {
			count++
		}
		if c.Point.x == x && c.Point.y == y-1 {
			count++
		}
		if c.Point.x == x+1 && c.Point.y == y-1 {
			count++
		}

		if c.Point.x == x-1 && c.Point.y == y+1 {
			count++
		}
		if c.Point.x == x && c.Point.y == y+1 {
			count++
		}
		if c.Point.x == x+1 && c.Point.y == y+1 {
			count++
		}

		if c.Point.x == x-1 && c.Point.y == y {
			count++
		}
		if c.Point.x == x+1 && c.Point.y == y {
			count++
		}
	}

	return count
}
