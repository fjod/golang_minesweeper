package golang_minesweeper

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

type MinePosition struct {
	Point     Point
	IsCleared bool
}

type Game struct {
	Mines []MinePosition
	Board *Board
	Steps int
}

func (g *Game) Refresh() {
	clearConsole()

}

func clearConsole() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Println("cant clear console")
	}
}
