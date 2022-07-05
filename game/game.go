package game

import (
	"errors"
	"fmt"
	"model/board"
	"model/player"
	"strconv"
)

type Game struct {
	Board    board.Board
	Player1  player.Player
	Player2  player.Player
	GridSize int
	turn     int
}

var newGame *Game

func CreateGame(Player1Name, Player2Name, Player1Mark, Player2Mark string, GridSize int) *Game {
	if newGame != nil {
		return newGame
	}
	newGame = CreateNewGame(Player1Name, Player2Name, Player1Mark, Player2Mark, GridSize)
	return newGame
}

func CreateNewGame(Player1Name, Player2Name, Player1Mark, Player2Mark string, GridSize int) *Game {

	var tempGame Game
	tempGame.Board = *board.CreateNewBoard(GridSize)
	tempGame.Player1 = *player.CreateNewPlayer(Player1Name, Player1Mark)
	tempGame.Player2 = *player.CreateNewPlayer(Player2Name, Player2Mark)
	tempGame.GridSize = GridSize
	tempGame.turn = 1

	return &tempGame
}

func (g *Game) PrintBoard() {
	n := 0
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

func (g *Game) PlayGame() bool {
	var MarkCellPlayerText string
	var MarkCellPlayer1 int
	var iswinbyplayer1 bool
ErrByP1:
	if g.turn > g.GridSize*g.GridSize {
		fmt.Println("Match is Draw")
		newGame = nil
		return true
	}
	fmt.Println("Player 1 turn")
	Gridmax := g.GridSize * g.GridSize

	fmt.Println("Select position to mark")
	fmt.Scanln(&MarkCellPlayerText)
	var err error
	MarkCellPlayer1, err = strconv.Atoi(MarkCellPlayerText)
	if MarkCellPlayer1 < 0 || MarkCellPlayer1 >= Gridmax || err != nil {
		fmt.Println("Enter Number In range")
		goto ErrByP1
	}

	isMarked := g.IsCellMArked(MarkCellPlayer1)
	if !isMarked {
		fmt.Println("Select position already marked")
		goto ErrByP1
	}

	iswinbyplayer1, err = g.Player1Turn(MarkCellPlayer1)
	if err != nil {
		fmt.Println(err)
	} else {
		if iswinbyplayer1 {
			newGame = nil
			return true
		}

	}

	var MarkCellPlayer2 int
ErrByP2:

	if g.turn > g.GridSize*g.GridSize {
		fmt.Println("Match is Draw")
		newGame = nil
		return true
	}
	fmt.Println("Player 2 turn")

	fmt.Println("Select position to mark")
	fmt.Scanln(&MarkCellPlayerText)

	MarkCellPlayer2, err = strconv.Atoi(MarkCellPlayerText)
	if MarkCellPlayer2 < 0 || MarkCellPlayer2 >= g.GridSize*g.GridSize || err != nil {
		fmt.Println("Enter Number In range")
		goto ErrByP2
	}
	isMarked = g.IsCellMArked(MarkCellPlayer2)
	if !isMarked {
		fmt.Println("Select position already marked")
		goto ErrByP2
	}

	iswinbyplayer2, err := g.Player2Turn(MarkCellPlayer2)
	if err != nil {
		fmt.Println(err)
	} else {

		if iswinbyplayer2 {
			newGame = nil
			return true
		}

	}
	return false
}
func (g *Game) Player1Turn(cellNo int) (bool, error) {
	//checkfor cell no if used return error
	if cellNo > g.GridSize*g.GridSize {

		return false, errors.New("not a valid cell")
	}

	err := g.MarkCell(cellNo, g.Player1)
	if err != nil {
		return false, err
	}
	g.PrintBoard()
	// Analuse result
	status := g.Board.Analyse(&g.Player1)
	g.turn += 1
	if status {
		newGame = nil
		return true, nil
	} else {

		return false, nil
	}
}
func (g *Game) Player2Turn(CellNo int) (bool, error) {
	//checkfor cell no if used return error
	if CellNo > g.GridSize*g.GridSize {

		return false, errors.New("not a valid cell")
	}
	fmt.Println(g.turn)

	//check for player turn
	err := g.MarkCell(CellNo, g.Player2)
	if err != nil {
		return false, err
	}
	g.PrintBoard()
	// Analuse result
	status := g.Board.Analyse(&g.Player2)
	g.turn += 1
	if status {
		newGame = nil
		return true, nil
	} else {

		return false, nil
	}
}
func (g *Game) MarkCell(cellNo int, Player player.Player) error {

	if g.Board.Cell[cellNo].Mark == strconv.Itoa(cellNo) {
		g.Board.Cell[cellNo].Mark = Player.Mark
		return nil
	}
	return errors.New(" Cell Already marked")

}
func (g *Game) IsCellMArked(position int) bool {

	return g.Board.Cell[position].Mark == strconv.Itoa(position)
}
