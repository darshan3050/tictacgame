package main

import (
	"fmt"
	"model/game"
)

func main() {

	var Player1Name, Player2Name, VarplayGame, Player1Mark, Player2Mark string
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
	if Player1Name == Player2Name {
		fmt.Println("Both Player Name cannot be Same")
		goto scanP2Name
	}
PlayerMark1:
	fmt.Println("Enter Player 1 Mark")
	_, err = fmt.Scan(&Player1Mark)
	if err != nil || len(Player1Mark) > 1 {
		fmt.Println("Enter a valid Name:", err)
		goto PlayerMark1
	}

PlayerMark2:
	fmt.Println("Enter Player 2 Mark")
	_, err = fmt.Scan(&Player2Mark)
	if err != nil || len(Player2Mark) > 1 {
		fmt.Println("Enter a valid Name:", err)
		goto PlayerMark2
	}
	if Player1Mark == Player2Mark {
		fmt.Println("Both Player Name cannot be Same")
		goto PlayerMark2
	}

getGrid:
	fmt.Println("Enter size of Gridbetween 3 to 7")
	_, err = fmt.Scan(&gridsize)
	if err != nil {
		fmt.Println("Enter a valid Name:", err)
		goto getGrid
	} else if gridsize < 3 || gridsize > 7 {
		fmt.Println("Enter a valid Range", err)
		goto getGrid
	}

	game := game.CreateGame(Player1Name, Player2Name, Player1Mark, Player2Mark, gridsize)

	game.PrintBoard()

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
