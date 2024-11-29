package game

import "fmt"

func isWinner(board [3][3]string) bool {
	lines := [][]string{
		// Rows
		{board[0][0], board[0][1], board[0][2]},
		{board[1][0], board[1][1], board[1][2]},
		{board[2][0], board[2][1], board[2][2]},
		// Columns
		{board[0][0], board[1][0], board[2][0]},
		{board[0][1], board[1][1], board[2][1]},
		{board[0][2], board[1][2], board[2][2]},
		// Diagonals
		{board[0][0], board[1][1], board[2][2]},
		{board[0][2], board[1][1], board[2][0]},
	}

	marks := []string{"o", "x"}

	for _, mark := range marks {
		for _, line := range lines {
			count := 0

			for _, cell := range line {
				if cell == mark {
					count++
				}
			}

			if count == 3 {
				return true
			}
		}
	}

	return false
}

func isDraw(board [3][3]string) bool {
	for _, row := range board {
		for _, cell := range row {
			if cell == "." {
				return false
			}
		}
	}

	return true
}

func isGameOver(board [3][3]string) bool {
	if isWinner(board) || isDraw(board) {
		return true
	}

	return false
}

func displayGameOverMessage(player int, board [3][3]string) {
	if isDraw(board) {
		fmt.Println("Draw!")
	} else {
		fmt.Printf("Player %v wins!\n", player)
	}
}
