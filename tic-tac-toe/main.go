package main

import (
	"fmt"
)

const (
	EMPTY    = iota // iota is used to create incrementing numbers starting from 0
	PLAYER_X        // PLAYER_X is represented by 1
	PLAYER_O
)

var board [3][3]int   // 3x3 board for the game
var currentPlayer int // keeps track of the current player

func main() {
	currentPlayer = PLAYER_X //PLAYER_X starts the game
	for {
		printBoard()    //print the current state of the board
		if checkWin() { // check if the current player has won
			fmt.Printf("Player %s wins!\n", playerSymbol(currentPlayer))
			break
		}
		if checkDraw() {
			// check if the game is draw
			fmt.Println("Its a draw!")
			break
		}
		playerMove()   // prompt the current player to make a move
		switchPlayer() // switch to the other player
	}
}

func printBoard() {
	fmt.Println("Current Board:")
	for i := 0; i < 3; i++ { // iterate over rows
		for j := 0; j < 3; j++ {
			fmt.Printf("%s ", playerSymbol(board[i][j])) // print the symbol for each cell
		}
		fmt.Println() //new line after each row
	}
}

func playerSymbol(player int) string {
	switch player {
	case PLAYER_X:
		return "X" // return "X" for PLAYER_X
	case PLAYER_O:
		return "O" // return "O" for PLAYER_O
	default:
		return "." // return "." for EMPTY cells
	}
}

func playerMove() {
	var row, col int
	for {
		fmt.Printf("Player %s, Enter your move (row and column):", playerSymbol(currentPlayer))
		_, err := fmt.Scanln(&row, &col) // read the row and column from the player
		if err != nil {
			fmt.Println("Invalid input. Please enter two integers.")
		}
		if row >= 0 && row < 3 && col >= 0 && col < 3 && board[row][col] == EMPTY {
			board[row][col] = currentPlayer // update the board with the player's move
			break
		} else {
			fmt.Println("Invalid move. Try again.") //
		}
	}
}

func switchPlayer() {
	if currentPlayer == PLAYER_X {
		currentPlayer = PLAYER_O // switch to PLAYER_O if currentPlayer is PLAYER_X
	} else {
		currentPlayer = PLAYER_X // switch to PLAYER_X if currentPlayer is PLAYER_O
	}
}

func checkWin() bool {
	for i := 0; i < 3; i++ {
		if board[i][0] == currentPlayer && board[i][1] == currentPlayer && board[i][2] == currentPlayer {
			fmt.Printf("Winning row: %d\n", i) // Debug print
			return true                        //check rows for a win
		}
		if board[0][i] == currentPlayer && board[1][i] == currentPlayer && board[2][i] == currentPlayer {
			fmt.Printf("Winning column: %d\n", i) // Debug print
			return true                           // check columns for a win
		}
	}
	if board[0][0] == currentPlayer && board[1][1] == currentPlayer && board[2][2] == currentPlayer {
		fmt.Println("Winning diagonal: top-left to bottom-right") // Debug print
		return true                                               //check diagonal from top-left to bottom-right
	}
	if board[0][2] == currentPlayer && board[1][1] == currentPlayer && board[2][0] == currentPlayer {
		fmt.Println("Winning diagonal: top-right to bottom-left") // Debug print
		return true                                               //check diagonal from top-right to bottom-left
	}
	return false
}

func checkDraw() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == EMPTY {
				return false // if any cell is empty, it's not a draw
			}
		}
	}
	return true
}
