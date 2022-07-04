package main

import (
	"fmt"
	"model/game"
)

func main() {

	var Player1Name, Player2Name, VarplayGame string
	var gridsize int
NewGame:
	isWin := false
scanP1Name:
	fmt.Println("Enter Player 1 Name")
	_, err := fmt.Scan(&Player1Name)
	if err != nil {
		fmt.Println("Enter a valid Name:", err)
		goto scanP1Name
	}
scanP2Name:
	fmt.Println("Enter Player 2 Name")
	_, err = fmt.Scan(&Player2Name)
	if err != nil {
		fmt.Println("Enter a valid Name:", err)
		goto scanP2Name
	}
getGrid:
	fmt.Println("Enter size of Grid")
	_, err = fmt.Scan(&gridsize)
	if err != nil {
		fmt.Println("Enter a valid Name:", err)
		goto getGrid
	} else if gridsize < 3 || gridsize > 7 {
		fmt.Println("Enter a valid Range")
		goto getGrid
	}

	game := game.CreateGame(Player1Name, Player2Name, gridsize)
	game.PrintBoard()
	fmt.Println(game.Board)

	for i := 0; !isWin; i++ {
		isWin = game.PlayGame()

	}

	fmt.Println("Press N for New Game: ")
	fmt.Scan(&VarplayGame)
	if VarplayGame == "n" || VarplayGame == "N" {
		game = nil
		fmt.Println(game)
		goto NewGame

	}

}
