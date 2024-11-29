package game

import "fmt"

func createBoard() [3][3]string {
	return [3][3]string{
		{".", ".", "."},
		{".", ".", "."},
		{".", ".", "."},
	}
}

func displayBoard(board [3][3]string) {
	for _, row := range board {
		fmt.Println(row)
	}

	fmt.Println()
}
