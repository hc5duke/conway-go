package main

import (
	"fmt"
	"time"
)

// Constants
var (
	Rows = 20
	Cols = 20
	Seed = "glider"
)

// boards
var (
	currentBoard = MakeBoard()
	nextBoard    = MakeBoard()
)

func main() {
	SeedBoard(currentBoard)
	// Clear board space
	for i := 0; i < Rows; i++ {
		fmt.Println("\n")
	}

	for i := 0; ; i++ {
		PrintBoard(currentBoard)
		Step(currentBoard, nextBoard)
		currentBoard = nextBoard
		nextBoard = MakeBoard()
		time.Sleep(100 * time.Millisecond)
	}
}
