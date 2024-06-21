package Game

import (
	"reflect"
	"testing"
)

func TestGetNearbyMinesCount(t *testing.T) {
	tests := []struct {
		name     string
		game     *Game
		x, y     int
		expected int
	}{
		{
			name: "No mines nearby",
			game: &Game{
				Mines: []MinePosition{},
			},
			x: 0, y: 0,
			expected: 0,
		},
		{
			name: "One mine nearby",
			game: &Game{
				Mines: []MinePosition{
					{Point: Point{x: 0, y: 1}},
				},
			},
			x: 0, y: 0,
			expected: 1,
		},
		{
			name: "Multiple mines nearby",
			game: &Game{
				Mines: []MinePosition{
					{Point: Point{x: 4, y: 4}},
					{Point: Point{x: 6, y: 6}},
					{Point: Point{x: 5, y: 4}},
				},
			},
			x: 5, y: 5,
			expected: 3,
		},
		{
			name: "No mines on board",
			game: &Game{
				Mines: []MinePosition{},
			},
			x: 2, y: 2,
			expected: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getNearbyMinesCount(tc.game, tc.x, tc.y)
			if actual != tc.expected {
				t.Errorf("get_nearby_mines_count(%v, %d, %d) = %d, expected %d", tc.game, tc.x, tc.y, actual, tc.expected)
			}
		})
	}
}

func c(x int, y int) CellCheckedState {
	return CellCheckedState{
		x: x,
		y: y,
	}
}

func TestGetNearbyCells(t *testing.T) {
	input := [][]CellCheckedState{
		{c(0, 0), c(1, 0), c(2, 0)},
		{c(0, 1), c(1, 1), c(2, 1)},
		{c(0, 2), c(1, 2), c(2, 2)},
	}
	eCenter := []CellCheckedState{
		c(0, 0), c(1, 0), c(2, 0),
		c(0, 1), c(2, 1),
		c(0, 2), c(1, 2), c(2, 2)}
	eTopLeft := []CellCheckedState{
		c(1, 0),
		c(0, 1), c(1, 1)}
	eBotRight := []CellCheckedState{
		c(1, 1), c(2, 1),
		c(1, 2)}

	inputChecked := [][]CellCheckedState{
		{CellCheckedState{
			x:       0,
			y:       0,
			checked: true,
		}, c(1, 0), c(2, 0)},
		{c(0, 1), c(1, 1), c(2, 1)},
		{c(0, 2), c(1, 2), c(2, 2)},
	}

	eCenterChecked := []CellCheckedState{
		c(1, 0), c(2, 0),
		c(0, 1), c(2, 1),
		c(0, 2), c(1, 2), c(2, 2)}

	tests := []struct {
		name     string
		s        [][]CellCheckedState
		x, y     int
		expected []CellCheckedState
	}{
		{
			name: "Get nearby cells for center cell",
			s:    input,
			x:    1, y: 1,
			expected: eCenter,
		},
		{
			name: "Get nearby cells for top-left corner cell",
			s:    input,
			x:    0, y: 0,
			expected: eTopLeft,
		},
		{
			name: "Get nearby cells for bottom-right corner cell",
			s:    input,
			x:    2, y: 2,
			expected: eBotRight,
		},
		{
			name: "Get nearby cells for center cell with one already checked",
			s:    inputChecked,
			x:    1, y: 1,
			expected: eCenterChecked,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getNearbyCells(&tc.s, tc.x, tc.y)
			for _, a := range actual {
				found := false
				for _, e := range tc.expected {
					if a.x == e.x && a.y == e.y {
						found = true
						break
					}
				}

				if !found {
					t.Errorf("get_nearby_cells(%v, %d, %d) = %v, expected %v", tc.s, tc.x, tc.y, actual, tc.expected)
				}
			}
		})
	}
}
func TestOpenNearbyCells(t *testing.T) {
	tests := []struct {
		name     string
		game     *Game
		expected [][]CellState
	}{
		{
			name: "Open cells on empty board",
			game: &Game{
				Board: &Board{
					Width:  3,
					Height: 3,
					Cells: [][]CellState{
						{Unknown, Unknown, Unknown},
						{Unknown, Unknown, Unknown},
						{Unknown, Unknown, Unknown},
					},
					current: Point{x: 1, y: 1},
				},
				Mines: []MinePosition{},
			},
			expected: [][]CellState{
				{Opened_no_mines_nearby, Opened_no_mines_nearby, Opened_no_mines_nearby},
				{Opened_no_mines_nearby, Opened_no_mines_nearby, Opened_no_mines_nearby},
				{Opened_no_mines_nearby, Opened_no_mines_nearby, Opened_no_mines_nearby},
			},
		},
		{
			name: "Open cells around mine",
			game: &Game{
				Board: &Board{
					Width:  3,
					Height: 3,
					Cells: [][]CellState{
						{Unknown, Unknown, Unknown},
						{Unknown, Unknown, Unknown},
						{Unknown, Unknown, Unknown},
					},
					current: Point{x: 1, y: 1},
				},
				Mines: []MinePosition{
					{Point: Point{x: 0, y: 0}},
				},
			},
			expected: [][]CellState{
				{Unknown, Unknown, Unknown},
				{Unknown, Opened_1_mine_nearby, Unknown},
				{Unknown, Unknown, Unknown},
			},
		},
		{
			name: "Open cells around mine",
			game: &Game{
				Board: &Board{
					Width:  3,
					Height: 3,
					Cells: [][]CellState{
						{Unknown, Unknown, Unknown},
						{Unknown, Unknown, Unknown},
						{Unknown, Unknown, Unknown},
					},
					current: Point{x: 2, y: 2},
				},
				Mines: []MinePosition{
					{Point: Point{x: 0, y: 0}},
				},
			},
			expected: [][]CellState{
				{Unknown, Opened_1_mine_nearby, Opened_no_mines_nearby},
				{Opened_1_mine_nearby, Opened_1_mine_nearby, Opened_no_mines_nearby},
				{Opened_no_mines_nearby, Opened_no_mines_nearby, Opened_no_mines_nearby},
			},
		},
		{
			name: "Open cells around multiple mines",
			game: &Game{
				Board: &Board{
					Width:  5,
					Height: 5,
					Cells: [][]CellState{
						{Unknown, Unknown, Unknown, Unknown, Unknown},
						{Unknown, Unknown, Unknown, Unknown, Unknown},
						{Unknown, Unknown, Unknown, Unknown, Unknown},
						{Unknown, Unknown, Unknown, Unknown, Unknown},
						{Unknown, Unknown, Unknown, Unknown, Unknown},
					},
					current: Point{x: 2, y: 2},
				},
				Mines: []MinePosition{
					{Point: Point{x: 0, y: 0}},
					{Point: Point{x: 4, y: 4}},
				},
			},
			expected: [][]CellState{
				{Unknown, Opened_1_mine_nearby, Opened_no_mines_nearby, Opened_no_mines_nearby, Opened_no_mines_nearby},
				{Opened_1_mine_nearby, Opened_1_mine_nearby, Opened_no_mines_nearby, Opened_no_mines_nearby, Opened_no_mines_nearby},
				{Opened_no_mines_nearby, Opened_no_mines_nearby, Opened_no_mines_nearby, Opened_no_mines_nearby, Opened_no_mines_nearby},
				{Opened_no_mines_nearby, Opened_no_mines_nearby, Opened_no_mines_nearby, Opened_1_mine_nearby, Opened_1_mine_nearby},
				{Opened_no_mines_nearby, Opened_no_mines_nearby, Opened_no_mines_nearby, Opened_1_mine_nearby, Unknown},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			openNearbyCells(tc.game)

			if !reflect.DeepEqual(tc.game.Board.Cells, tc.expected) {
				t.Errorf("openNearbyCells = %v, expected %v", (*tc.game).Board.Cells, tc.expected)
			}
		})
	}
}
