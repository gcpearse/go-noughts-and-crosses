package game

import "fmt"

func createBoard() [3][3]string {
	var board = [3][3]string{}

	for i, row := range board {
		for j := range row {
			board[i][j] = "."
		}
	}

	return board
}

func displayBoard(board [3][3]string) {
	for _, row := range board {
		fmt.Println(row)
	}

	fmt.Println()
}
