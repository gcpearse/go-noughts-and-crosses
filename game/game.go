package game

import "fmt"

func StartGame() {
	board := createBoard()

	player := 2

	for !isGameOver(board) {
		if player == 1 {
			player = 2
		} else if player == 2 {
			player = 1
		}

		turn(player, &board)
	}

	if isGameOver(board) {
		displayGameOverMessage(player, board)

		fmt.Println()

		displayBoard(board)
	}
}
