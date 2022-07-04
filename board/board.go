package board

import (
	"fmt"
	"model/cell"
	"model/player"
)

type Board struct {
	Cell []cell.Cell
	Size int
}

var newBoard *Board

func CreateNewBoard(BoardSize int) *Board {
	if newBoard != nil {
		return newBoard
	} else {
		newBoard = CreateBoard(BoardSize)
		return newBoard
	}
}
func CreateBoard(BoardSize int) *Board {
	var TempBoard Board
	for i := 0; i < BoardSize*BoardSize; i++ {
		TempBoard.Cell = append(TempBoard.Cell, *cell.NewCell())

	}
	TempBoard.Size = BoardSize
	return &TempBoard
}

func (b *Board) Analyse(p *player.Player) bool {
	status := false // check all verticals
	status = b.AnalyseVertical(p)
	if status {
		fmt.Println(p.Name, " Won the round by Vertical Match")
		return status
	}
	//check all horizontal
	status = b.AnalyseHorizontal(p)
	if status {
		fmt.Println(p.Name, " Won the round by Horizontal Match")
		return status
	}
	//check all diagonal
	status = b.AnalyseDiagonal(p)
	if status {
		fmt.Println(p.Name, " Won the round by Diagonal Match")
		return status
	}
	return status
}
func (b *Board) AnalyseVertical(p *player.Player) bool {
	var count int = 0

	for i := 0; i < b.Size; i++ {
		for j := 0; j < b.Size*b.Size; j += b.Size {

			if b.Cell[j+i].Mark == p.Mark && b.Cell[j+i].Mark != "" {

				count = count + 1

			}
		}
		if count == b.Size {
			return true
		} else {
			count = 0
		}

	}

	return false

}
func (b *Board) AnalyseHorizontal(p *player.Player) bool {
	var count int = 0
	for i := 0; i < b.Size*b.Size; i += b.Size {
		for j := 0; j < b.Size; j++ {
			if b.Cell[i+j].Mark == p.Mark && b.Cell[i+j].Mark != "" {
				count = count + 1
			}
		}
		if count == b.Size {
			return true
		} else {
			count = 0
		}

	}

	return false

}

func (b *Board) AnalyseDiagonal(p *player.Player) bool {
	var countleft int = 0
	var countright int = 0
	for i := 0; i < b.Size*b.Size; i += b.Size + 1 {
		if b.Cell[i].Mark == p.Mark && b.Cell[i].Mark != "" {
			countleft = countleft + 1
		}
	}

	for i := b.Size - 1; i < b.Size*b.Size; i += b.Size - 1 {
		if b.Cell[i].Mark == p.Mark && b.Cell[i].Mark != "" {
			countright = countright + 1
		}
	}
	if countleft == b.Size || countright == b.Size {
		return true
	} else {
		return false
	}

}
