package Game

type CellCheckedState struct {
	checked bool
	x       int
	y       int
	state   CellState
}

func open_nearby_cells(g *Game) {
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
}

func get_nearby_cells(s *[][]CellCheckedState, x int, y int) []CellCheckedState {
	ret := make([]CellCheckedState, 8)
	w := len(*s)
	h := len((*s)[0])
	append_cell_if_possible(&ret, s, x-1, y-1, w, h)
	append_cell_if_possible(&ret, s, x, y-1, w, h)
	append_cell_if_possible(&ret, s, x+1, y-1, w, h)
	append_cell_if_possible(&ret, s, x-1, y+1, w, h)
	append_cell_if_possible(&ret, s, x, y+1, w, h)
	append_cell_if_possible(&ret, s, x+1, y+1, w, h)
	append_cell_if_possible(&ret, s, x+1, y, w, h)
	append_cell_if_possible(&ret, s, x-1, y, w, h)
	return ret
}

func append_cell_if_possible(ret *[]CellCheckedState, s *[][]CellCheckedState, x int, y int, w int, h int) {
	if x >= 0 && x < w && y >= 0 && y < h {
		*ret = append(*ret, (*s)[x][y])
	}
}

func get_nearby_mines_count(g *Game, x int, y int) int {
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
