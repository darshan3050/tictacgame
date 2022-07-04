package game

import (
	"fmt"
	"model/board"
	"testing"
)

var testdataforPrintingBoard = []struct {
	GridSize           int
	expectearrayLength int
}{
	{3, 9}, {5, 25}, {7, 49}, {4, 16}, {6, 36},
}

func TestPrintBoard(t *testing.T) {
	var g Game
	for o := 0; o < len(testdataforPrintingBoard); o++ {
		// boardsize := testdataforPrintingBoard[o].GridSize
		g.Board = *board.CreateBoard(testdataforPrintingBoard[o].GridSize)
		n := 0
		for i := 0; i < g.GridSize; i++ {
			for j := 0; j < g.GridSize; j++ {
				fmt.Print(g.Board.Cell[n])
				n = n + 1
			}
			fmt.Println()
		}
	}
}
