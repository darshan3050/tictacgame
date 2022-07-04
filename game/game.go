package game

import (
	"errors"
	"fmt"
	"model/board"
	"model/player"
)

type Game struct {
	Board    board.Board
	Player1  player.Player
	Player2  player.Player
	GridSize int
	turn     int
}

var newGame *Game

func CreateGame(Player1Name, Player2Name string, GridSize int) *Game {
	if newGame != nil {
		return newGame
	}
	newGame = CreateNewGame(Player1Name, Player2Name, GridSize)
	return newGame
}

func CreateNewGame(Player1Name, Player2Name string, GridSize int) *Game {

	var tempGame Game
	tempGame.Board = *board.CreateNewBoard(GridSize)
	tempGame.Player1 = *player.CreateNewPlayer(Player1Name, "X")
	tempGame.Player2 = *player.CreateNewPlayer(Player2Name, "O")
	tempGame.GridSize = GridSize
	tempGame.turn = 0

	return &tempGame
}

func (g *Game) PrintBoard() {
	n := 0
	for i := 0; i < g.GridSize; i++ {
		for j := 0; j < g.GridSize; j++ {
			fmt.Printf("%s ", g.Board.Cell[n])
			n = n + 1
		}
		fmt.Println()
	}
}

func (g *Game) PlayGame() bool {
	var MarkCellPlayer1 int
	fmt.Println("Player 1 turn")
	fmt.Println("Select position to mark")
	fmt.Scan(&MarkCellPlayer1)
	iswinbyplayer1, err := g.Player1Turn(MarkCellPlayer1)
	if err != nil {
		fmt.Println(err)
	} else {
		if iswinbyplayer1 {
			return true
		}

	}

	var MarkCellPlayer2 int
	fmt.Println("Player 2 turn")
	fmt.Println("Select position to mark")
	fmt.Scan(&MarkCellPlayer2)
	iswinbyplayer2, err := g.Player2Turn(MarkCellPlayer2)
	if err != nil {
		fmt.Println(err)
	} else {
		if iswinbyplayer2 {
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

	// ceeck for turns >8 deaw
	if g.turn > g.GridSize*g.GridSize {
		fmt.Println("Match is Draw")
		newGame = nil
		return true, nil
	}

	//check for player turn
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

	// ceeck for turns >8 deaw
	if g.turn > g.GridSize*g.GridSize {
		fmt.Println("Match is Draw")
		newGame = nil
		return true, nil
	}

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

	fmt.Println(" mark cell index ", cellNo)
	if g.Board.Cell[cellNo].Mark != "" {
		return errors.New(" Cell Already marked")
	}
	g.Board.Cell[cellNo].Mark = Player.Mark

	return nil
}
func (g *Game) ClearBoard() {

	n := 0
	for i := 0; i < g.GridSize; i++ {
		for j := 0; j < g.GridSize; j++ {
			g.Board.Cell[n].Mark = ""
			n = n + 1
		}
		fmt.Println()
	}
}
