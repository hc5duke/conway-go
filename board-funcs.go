package main

import (
	"fmt"
)

func MakeBoard() [][]bool {
	_board := make([][]bool, Rows)
	for i := range _board {
		_board[i] = make([]bool, Cols)
		for j := range _board[i] {
			_board[i][j] = false
		}
	}
	return _board
}

func PrintBoard(board [][]bool) {
	str := fmt.Sprintf("[%dA", Rows+3)
	fmt.Println("\x1B" + str)
	var borderHoriz = "═"
	for i := 0; i < Cols; i++ {
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

func Step(curr, next [][]bool) {
	for i, row := range curr {
		for j, cell := range row {
			aliveNeighbors := CountAliveNeighbors(curr, i, j)
			if cell {
				//Any live cell with fewer than two live neighbors dies, as if caused by under-population.
				//Any live cell with two or three live neighbors lives on to the next generation.
				//Any live cell with more than three live neighbors dies, as if by overcrowding.
				next[i][j] = aliveNeighbors == 2 || aliveNeighbors == 3
			} else {
				//Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.
				next[i][j] = aliveNeighbors == 3
			}
		}
	}
}

func CountAliveNeighbors(board [][]bool, row, col int) int {
	count := 0
	for i := row - 1; i <= row+1; i++ {
		for j := col - 1; j <= col+1; j++ {
			if (i == row && j == col) || // center
				i < 0 || j < 0 || // top/left
				i >= Rows || j >= Cols || // bottom/right
				!board[i][j] { // dead
				continue
			}
			count++ // alive neighbor
		}
	}
	return count
}
