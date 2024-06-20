package Game

type CellCheckedState struct {
	checked bool
	x       int
	y       int
	state   CellState
}

func openNearbyCells(g *Game) {
	// сначала создать массив клеток с начальным состоянием
	h := g.Board.height
	w := g.Board.width
	state := make([][]CellCheckedState, h)
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

	openNearbyCellsRec(g, &state, g.Board.current.x, g.Board.current.y)
}

func openNearbyCellsRec(g *Game, s *[][]CellCheckedState, x int, y int) {
	// проверяем выбранную клетку
	(*s)[x][y].checked = true
	minesAround := getNearbyMinesCount(g, x, y)
	if minesAround > 0 {
		g.Board.cells[x][y] = FromCount(minesAround) // установить кол-во мин в клетке
		return                                       // если рядом с клеткой мины, то останавливаем поиск соседних клеток
	}
	if minesAround == 0 {
		g.Board.cells[x][y] = Opened_no_mines_nearby // открыть клетку
	}

	surround := getNearbyCells(s, x, y) // ищем всех соседей клетки

	for _, cell := range surround {
		openNearbyCellsRec(g, s, cell.x, cell.y)
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
