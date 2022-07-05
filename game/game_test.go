package game

import (
	"fmt"
	"model/board"
	"strconv"
	"testing"
)

var testdataforPrintingBoard = []struct {
	GridSize int
}{
	{3}, {5}, {7}, {4}, {6},
}
var testdataformarkcell = []struct {
	input    int
	output   string
	expected bool
}{
	{3, "3", true}, {5, "5", true}, {7, "X", false}, {4, "O", false}, {16, "16", true},
}

func TestPrintBoard(t *testing.T) {
	// n := 0
	var g Game
	for i := 0; i < len(testdataforPrintingBoard); i++ {
		n := 0
		g.GridSize = testdataforPrintingBoard[i].GridSize
		g.Board = *board.CreateNewBoard(testdataforPrintingBoard[i].GridSize)
		for i := 0; i < g.GridSize; i++ {
			for j := 0; j < g.GridSize; j++ {
				if n < 10 {
					fmt.Print("|   ", g.Board.Cell[n].Mark, "  |")
					n = n + 1
				} else {
					fmt.Print("|  ", g.Board.Cell[n].Mark, "  |")
					n = n + 1
				}
			}
			fmt.Print("\n")
			for k := 0; k < g.GridSize; k++ {

				fmt.Print("========")
			}
			fmt.Print("\n")

		}
	}
}

func TestIsCellMArked(t *testing.T) {
	for i := 0; i < len(testdataformarkcell); i++ {
		if testdataformarkcell[i].output == strconv.Itoa(testdataformarkcell[i].input) {
			fmt.Println("true Condition  ", "expected Bool is:", testdataformarkcell[i].expected, "expected board data was: ", testdataformarkcell[i].output, "test data was: ", strconv.Itoa(testdataformarkcell[i].input))
		} else {
			fmt.Println("false Condition  ", "expected Bool is:", testdataformarkcell[i].expected, "expected board data was: ", testdataformarkcell[i].output, "test data was: ", strconv.Itoa(testdataformarkcell[i].input))
		}
	}

}
