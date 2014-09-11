package main

import "fmt"

var (
	rows         = 20
	cols         = 20
	currentBoard = MakeBoard(rows, cols)
	nextBoard    = MakeBoard(rows, cols)
)

func MakeBoard(rows, cols int) [][]bool {
	_board := make([][]bool, rows)
	for i := range _board {
		_board[i] = make([]bool, cols) /* again the type? */
		for j := range _board[i] {
			_board[i][j] = false
		}
	}
	return _board
}

func PrintBoard(board [][]bool) {
	//console.log "\x1B" + "[#{ROWS + 3}A"
	var borderHoriz = "═"
	for i := 0; i < cols; i++ {
		borderHoriz += "══"
	}
	fmt.Println("╔" + borderHoriz + "╗")
	for i := range board {
		fmt.Print("║ ")
		for j := range board[i] {
			if board[i][j] {
				fmt.Print("█ ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println("║")
	}
	fmt.Println("╚" + borderHoriz + "╝")
}

func main() {
	PrintBoard(currentBoard)
}
