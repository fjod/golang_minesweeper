package main

import (
	"testing"
)

func TestCellStateString(t *testing.T) {
	tests := []struct {
		name      string
		cellState CellState
		expected  rune
	}{
		{
			name:      "Empty cell",
			cellState: Unknown,
			expected:  '\u25A0',
		},
		{
			name:      "Opened_no_mines_nearby cell",
			cellState: Opened_no_mines_nearby,
			expected:  '\u25A1',
		},
		{
			name:      "Mine cell",
			cellState: Mine,
			expected:  '\U0001F7E5',
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.cellState.GetRune()
			t.Logf("CellState.String() = %q, expected %q", actual, tc.expected)
			if actual != tc.expected {
				t.Errorf("CellState.String() = %q, expected %q", actual, tc.expected)
			}
		})
	}
}
