package game

import "fmt"

func StartGame() {
	board := createBoard()

	player := 2

	for !isGameOver(board) {
		switch player {
		case 1:
			player = 2
		case 2:
			player = 1
		}

		turn(player, &board)
	}

	if isGameOver(board) {
		displayBoard(board)

		displayGameOverMessage(player, board)

		fmt.Println()
	}
}
