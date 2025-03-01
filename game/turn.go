package game

import "fmt"

func turn(player int, board *[3][3]string) {
	var mark string

	if player == 1 {
		mark = "o"
	} else if player == 2 {
		mark = "x"
	}

	fmt.Printf("Player %v's turn. Mark a space with '%v'.\n", player, mark)

	fmt.Println()

	displayBoard(*board)

	var x, y int

	for {
		x, y = 0, 0

		for x < 1 || x > 3 {
			fmt.Print("Row: ")
			fmt.Scan(&x)
		}

		for y < 1 || y > 3 {
			fmt.Print("Column: ")
			fmt.Scan(&y)
		}

		if board[x-1][y-1] == "." {
			fmt.Println()
			break
		} else {
			fmt.Println()
			fmt.Println("Mark an empty space.")
			fmt.Println()
		}
	}

	board[x-1][y-1] = mark
}
