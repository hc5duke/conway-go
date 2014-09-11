package main

import "fmt"

var (
	rows  = 20
	cols  = 20
	board = MakeBoard(rows, cols)
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
func main() {
	fmt.Println("xxx")
	fmt.Println(board)
}
