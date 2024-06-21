package Game

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
)

func FromCount(c int) CellState {
	if c < 1 || c > 9 {
		return Unknown
	}
	return CellState(c + 2)
}
