package Game

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
func TestFromCount(t *testing.T) {
	tests := []struct {
		name     string
		count    int
		expected CellState
	}{
		{
			name:     "Count 0",
			count:    0,
			expected: Unknown,
		},
		{
			name:     "Count 1",
			count:    1,
			expected: Opened_1_mine_nearby,
		},
		{
			name:     "Count 9",
			count:    9,
			expected: Opened_9_mine_nearby,
		},
		{
			name:     "Count 10",
			count:    10,
			expected: Unknown,
		},
		{
			name:     "Count -1",
			count:    -1,
			expected: Unknown,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := FromCount(tc.count)
			if actual != tc.expected {
				t.Errorf("FromCount(%d) = %v, expected %v", tc.count, actual, tc.expected)
			}
		})
	}
}
