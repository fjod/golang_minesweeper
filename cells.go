package golang_minesweeper

type CellState int

const (
	Unknown CellState = iota
	Opened_no_mines_nearby
	Mine
	Opened_1_mine_nearby
	Opened_2_mine_nearby
	Opened_3_mine_nearby
	Opened_4_mine_nearby
	Opened_5_mine_nearby
	Opened_6_mine_nearby
	Opened_7_mine_nearby
	Opened_8_mine_nearby
	Opened_9_mine_nearby
	Flagged
	Selected_cell
)

var cellStateImage = map[CellState]rune{
	Unknown:                '\u25A0',     // black square
	Opened_no_mines_nearby: '\u25A1',     // white square
	Mine:                   '\U0001F7E5', // large red square
	Opened_1_mine_nearby:   '\u0031',
	Opened_2_mine_nearby:   '\u0032',
	Opened_3_mine_nearby:   '\u0033',
	Opened_4_mine_nearby:   '\u0034',
	Opened_5_mine_nearby:   '\u0035',
	Opened_6_mine_nearby:   '\u0036',
	Opened_7_mine_nearby:   '\u0037',
	Opened_8_mine_nearby:   '\u0038',
	Opened_9_mine_nearby:   '\u0039',
	Flagged:                '\u2691', // flag
	Selected_cell:          '\u25C9', // fisheye
}

func (cellState CellState) GetRune() rune {
	return cellStateImage[cellState]
}
