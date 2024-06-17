package golang_minesweeper

import (
	"reflect"
	"testing"
)

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
				width:   3,
				height:  3,
				current: Point{x: 1, y: 0},
			},
			move: Up,
			expected: &Board{
				width:   3,
				height:  3,
				current: Point{x: 1, y: 0},
			},
		},
		{
			name: "Move down from bottom row",
			board: &Board{
				width:   3,
				height:  3,
				current: Point{x: 1, y: 2},
			},
			move: Down,
			expected: &Board{
				width:   3,
				height:  3,
				current: Point{x: 1, y: 2},
			},
		},
		{
			name: "Move left from left column",
			board: &Board{
				width:   3,
				height:  3,
				current: Point{x: 0, y: 1},
			},
			move: Left,
			expected: &Board{
				width:   3,
				height:  3,
				current: Point{x: 0, y: 1},
			},
		},
		{
			name: "Move right from right column",
			board: &Board{
				width:   3,
				height:  3,
				current: Point{x: 2, y: 1},
			},
			move: Right,
			expected: &Board{
				width:   3,
				height:  3,
				current: Point{x: 2, y: 1},
			},
		},
		{
			name: "Move up from middle",
			board: &Board{
				width:   3,
				height:  3,
				current: Point{x: 1, y: 1},
			},
			move: Up,
			expected: &Board{
				width:   3,
				height:  3,
				current: Point{x: 1, y: 0},
			},
		},
		{
			name: "Move down from middle",
			board: &Board{
				width:   3,
				height:  3,
				current: Point{x: 1, y: 1},
			},
			move: Down,
			expected: &Board{
				width:   3,
				height:  3,
				current: Point{x: 1, y: 2},
			},
		},
		{
			name: "Move left from middle",
			board: &Board{
				width:   3,
				height:  3,
				current: Point{x: 1, y: 1},
			},
			move: Left,
			expected: &Board{
				width:   3,
				height:  3,
				current: Point{x: 0, y: 1},
			},
		},
		{
			name: "Move right from middle",
			board: &Board{
				width:   3,
				height:  3,
				current: Point{x: 1, y: 1},
			},
			move: Right,
			expected: &Board{
				width:   3,
				height:  3,
				current: Point{x: 2, y: 1},
			},
		},
		{
			name: "Move none",
			board: &Board{
				width:   3,
				height:  3,
				current: Point{x: 1, y: 1},
			},
			move: None,
			expected: &Board{
				width:   3,
				height:  3,
				current: Point{x: 1, y: 1},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.board.Move(tc.move)
			if !reflect.DeepEqual(tc.board, tc.expected) {
				t.Errorf("Board.Move(%v) = %v, expected %v", tc.move, tc.board, tc.expected)
			}
		})
	}
}
