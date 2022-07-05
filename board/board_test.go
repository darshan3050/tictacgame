package board

import (
	"fmt"
	"model/cell"
	"testing"
)

var testdataforvertical = []struct {
	FirstCellIndex  int
	SecondCellIndex int
	ThirdCellIndex  int
	ExpextedCount   int
}{
	{0, 3, 6, 3}, {1, 4, 7, 3}, {2, 5, 8, 3}, {0, 4, 8, 1},
}
var testdataforhorizontal = []struct {
	FirstCellIndex  int
	SecondCellIndex int
	ThirdCellIndex  int
	ExpextedCount   int
}{
	{0, 1, 2, 3}, {3, 4, 5, 3}, {6, 7, 8, 3}, {0, 3, 8, 1},
}
var testdatafordiagonal = []struct {
	FirstCellIndex  int
	SecondCellIndex int
	ThirdCellIndex  int
	ExpextedCount   int
}{
	{0, 4, 8, 3}, {2, 4, 6, 3}, {0, 2, 8, 2}, {0, 4, 2, 2},
}

func TestCreateBoard(t *testing.T) {
	BoardSize := 3
	var TempBoard Board
	for i := 0; i < BoardSize*BoardSize; i++ {
		TempBoard.Cell = append(TempBoard.Cell, *cell.NewCell(i))

	}
	TempBoard.Size = BoardSize
	fmt.Println("Create Board Test", TempBoard)

}

func TestAnalyseVertical(t *testing.T) {
	fmt.Println("Vertical Analysis task")
	var boardsize = 3
	var countout int = 0
	b := CreateBoard(boardsize)
	for o := 0; o < len(testdataforvertical); o++ {

		b.Cell[testdataforvertical[o].FirstCellIndex].Mark = "X"
		b.Cell[testdataforvertical[o].SecondCellIndex].Mark = "X"
		b.Cell[testdataforvertical[o].ThirdCellIndex].Mark = "X"
		for i := 0; i < boardsize; i++ {
			count := 0
			for j := 0; j < boardsize*boardsize; j += boardsize {

				if b.Cell[j+i].Mark == "X" && b.Cell[j+i].Mark != "" {
					countout = count + 1
					count++
				}
			}

		}
		if countout == testdataforvertical[o].ExpextedCount {

			fmt.Println("Testcase  no", o, "Actual Count: ", countout, "Expected Count:", testdataforvertical[o].ExpextedCount)
		} else {
			fmt.Println("Failed Testcase  no", o, "Actual Count: ", countout, "Expected Count:", testdataforhorizontal[o].ExpextedCount)
		}
		b.Cell[testdataforvertical[o].FirstCellIndex].Mark = ""
		b.Cell[testdataforvertical[o].SecondCellIndex].Mark = ""
		b.Cell[testdataforvertical[o].ThirdCellIndex].Mark = ""
	}

}
func TestAnalyseHorizontal(t *testing.T) {
	fmt.Println("Horizontal Analysis Task")
	var boardsize = 3
	var countout int = 0
	b := CreateBoard(boardsize)
	for o := 0; o < len(testdataforhorizontal); o++ {

		b.Cell[testdataforhorizontal[o].FirstCellIndex].Mark = "O"
		b.Cell[testdataforhorizontal[o].SecondCellIndex].Mark = "O"
		b.Cell[testdataforhorizontal[o].ThirdCellIndex].Mark = "O"

		for i := 0; i < b.Size*b.Size; i += b.Size {
			count := 0
			for j := 0; j < b.Size; j++ {
				if b.Cell[i+j].Mark == "O" && b.Cell[i+j].Mark != "" {

					countout = count + 1
					count = count + 1
				}
			}

		}
		if countout == testdataforhorizontal[o].ExpextedCount {

			fmt.Println("Testcase  no", o, "Actual Count: ", countout, "Expected Count:", testdataforhorizontal[o].ExpextedCount)
		} else {
			fmt.Println("Failed Testcase  no", o, "Actual Count: ", countout, "Expected Count:", testdataforhorizontal[o].ExpextedCount)
		}
		b.Cell[testdataforhorizontal[o].FirstCellIndex].Mark = ""
		b.Cell[testdataforhorizontal[o].SecondCellIndex].Mark = ""
		b.Cell[testdataforhorizontal[o].ThirdCellIndex].Mark = ""
	}
}
func TestAnalyseDiagonal(t *testing.T) {
	fmt.Println("Diagonal Analysis Task")
	var boardsize = 3

	b := CreateBoard(boardsize)
	for o := 0; o < len(testdatafordiagonal); o++ {
		countoutleft, countoutright := 0, 0
		countleft, countright := 0, 0
		b.Cell[testdatafordiagonal[o].FirstCellIndex].Mark = "O"
		b.Cell[testdatafordiagonal[o].SecondCellIndex].Mark = "O"
		b.Cell[testdatafordiagonal[o].ThirdCellIndex].Mark = "O"
		if o%2 == 0 {
			for i := 0; i < b.Size*b.Size; i += b.Size + 1 {
				if b.Cell[i].Mark == "O" && b.Cell[i].Mark != "" {
					countoutleft = countoutleft + 1
					countleft = countleft + 1
				}
			}
			if countoutleft == testdatafordiagonal[o].ExpextedCount {

				fmt.Println("Testcase  no", o, "Actual Count: ", countoutleft, "Expected Count:", testdatafordiagonal[o].ExpextedCount)
			} else {
				fmt.Println("Failed Testcase  no", o, "Actual Count: ", countoutleft, "Expected Count:", testdatafordiagonal[o].ExpextedCount)
			}

		} else {
			for i := b.Size - 1; i < b.Size*b.Size; i += b.Size - 1 {
				if b.Cell[i].Mark == "O" && b.Cell[i].Mark != "" {
					countoutright = countright + 1
					countright = countright + 1
				}
			}
			if countoutright == testdatafordiagonal[o].ExpextedCount {

				fmt.Println("Testcase  no", o, "Actual Count: ", countoutright, "Expected Count:", testdatafordiagonal[o].ExpextedCount)
			} else {
				fmt.Println("Failed Testcase  no", o, "Actual Count: ", countoutright, "Expected Count:", testdatafordiagonal[o].ExpextedCount)
			}
		}

		b.Cell[testdatafordiagonal[o].FirstCellIndex].Mark = ""
		b.Cell[testdatafordiagonal[o].SecondCellIndex].Mark = ""
		b.Cell[testdatafordiagonal[o].ThirdCellIndex].Mark = ""
	}

}
