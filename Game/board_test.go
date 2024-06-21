package Game

import (
	"reflect"
	"testing"
)

func TestMoveState(t *testing.T) {
	tests := []struct {
		name string
		from CellState
		move Direction
		to   CellState
	}{
		{
			name: "Move up from top row",
			from: Unknown,
			move: Down,
			to:   Selected_cell,
		},
		{
			name: "Move up from top row",
			from: Flagged,
			move: Right,
			to:   Selected_cell,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var board Board
			board.Height = 3
			board.Width = 3
			board.current = Point{x: 0, y: 0}
			board.Init()
			board.prevState = tc.from
			board.Move(tc.move)
			if !reflect.DeepEqual(board.Cells[board.current.x][board.current.y], tc.to) {
				t.Errorf("Board.Move(%v) = %v, expected %v", tc.move, board, tc.to)
			}
			if !reflect.DeepEqual(board.Cells[0][0], tc.from) {
				t.Errorf("Board.Move(%v) = %v, expected %v", tc.move, board, tc.to)
			}
		})
	}
}

func TestBoardMove(t *testing.T) {
	tests := []struct {
		name     string
		board    *Board
		move     Direction
		expected *Board
	}{
		{
			name: "Move up from top row",
			board: &Board{
				Width:   3,
				Height:  3,
				current: Point{x: 1, y: 0},
			},
			move: Up,
			expected: &Board{
				Width:   3,
				Height:  3,
				current: Point{x: 1, y: 0},
			},
		},
		{
			name: "Move down from bottom row",
			board: &Board{
				Width:   3,
				Height:  3,
				current: Point{x: 1, y: 2},
			},
			move: Down,
			expected: &Board{
				Width:   3,
				Height:  3,
				current: Point{x: 1, y: 2},
			},
		},
		{
			name: "Move left from left column",
			board: &Board{
				Width:   3,
				Height:  3,
				current: Point{x: 0, y: 1},
			},
			move: Left,
			expected: &Board{
				Width:   3,
				Height:  3,
				current: Point{x: 0, y: 1},
			},
		},
		{
			name: "Move right from right column",
			board: &Board{
				Width:   3,
				Height:  3,
				current: Point{x: 2, y: 1},
			},
			move: Right,
			expected: &Board{
				Width:   3,
				Height:  3,
				current: Point{x: 2, y: 1},
			},
		},
		{
			name: "Move up from middle",
			board: &Board{
				Width:   3,
				Height:  3,
				current: Point{x: 1, y: 1},
			},
			move: Up,
			expected: &Board{
				Width:   3,
				Height:  3,
				current: Point{x: 1, y: 0},
			},
		},
		{
			name: "Move down from middle",
			board: &Board{
				Width:   3,
				Height:  3,
				current: Point{x: 1, y: 1},
			},
			move: Down,
			expected: &Board{
				Width:   3,
				Height:  3,
				current: Point{x: 1, y: 2},
			},
		},
		{
			name: "Move left from middle",
			board: &Board{
				Width:   3,
				Height:  3,
				current: Point{x: 1, y: 1},
			},
			move: Left,
			expected: &Board{
				Width:   3,
				Height:  3,
				current: Point{x: 0, y: 1},
			},
		},
		{
			name: "Move right from middle",
			board: &Board{
				Width:   3,
				Height:  3,
				current: Point{x: 1, y: 1},
			},
			move: Right,
			expected: &Board{
				Width:   3,
				Height:  3,
				current: Point{x: 2, y: 1},
			},
		},
		{
			name: "Move none",
			board: &Board{
				Width:   3,
				Height:  3,
				current: Point{x: 1, y: 1},
			},
			move: None,
			expected: &Board{
				Width:   3,
				Height:  3,
				current: Point{x: 1, y: 1},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.board.Init()
			tc.board.Move(tc.move)
			if !reflect.DeepEqual(tc.board.current, tc.expected.current) {
				t.Errorf("Board.Move(%v) = %v, expected %v", tc.move, tc.board, tc.expected)
			}
		})
	}
}
