package main

import (
	"testing"
)

func TestInit(t *testing.T) {
	g := &Game{}
	g.Init()

	if g.Board == nil {
		t.Errorf("Game.Init() failed to initialize the Board")
	}

	if g.Board.height != 10 || g.Board.width != 10 {
		t.Errorf("Game.Init() initialized Board with incorrect dimensions")
	}

	if len(g.Mines) != (g.Board.height*g.Board.width)/10 {
		t.Errorf("Game.Init() generated incorrect number of mines")
	}
}

func TestGenerateRandomMines(t *testing.T) {
	g := &Game{
		Board: &Board{
			height: 10,
			width:  10,
		},
	}

	generateRandomMines(g)

	if len(g.Mines) != (g.Board.height*g.Board.width)/10 {
		t.Errorf("generateRandomMines() generated incorrect number of mines")
	}
}

func TestGenerateUniqueMines(t *testing.T) {
	g := &Game{
		Board: &Board{
			height: 10,
			width:  10,
		},
		Mines: []MinePosition{
			{Point{0, 0}, false},
			{Point{1, 1}, false},
		},
	}

	generateUniqueMines(g)

	if len(g.Mines) != 3 {
		t.Errorf("generateUniqueMines() generated duplicate mine position")
	}

	for i := 0; i < len(g.Mines)-1; i++ {
		for j := i + 1; j < len(g.Mines); j++ {
			if g.Mines[i].Point == g.Mines[j].Point {
				t.Errorf("generateUniqueMines() generated duplicate mine position")
			}
		}
	}
}
