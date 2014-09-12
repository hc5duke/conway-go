package main

import (
	"math/rand"
)

func SeedBoard(board [][]bool) {
	var seeds [][]int
	switch Seed {
	default:
	case "blank":
		seeds = make([][]int, 0)
	case "blinker":
		seeds = [][]int{
			[]int{2, 1},
			[]int{2, 2},
			[]int{2, 3},
		}
	case "toad":
		seeds = [][]int{
			[]int{2, 2},
			[]int{2, 3},
			[]int{2, 4},
			[]int{3, 1},
			[]int{3, 2},
			[]int{3, 3},
		}
	case "beacon":
		seeds = [][]int{
			[]int{1, 1},
			[]int{1, 2},
			[]int{2, 1},
			[]int{3, 4},
			[]int{4, 3},
			[]int{4, 4},
		}
	case "pulsar":
		seeds = make([][]int, 48)
		index := 0
		for _, num1 := range []int{2, 7, 9, 14} {
			for _, num2 := range []int{4, 5, 6, 10, 11, 12} {
				seeds[index] = []int{num1, num2}
				seeds[index+1] = []int{num2, num1}
				index += 2
			}
		}
	// Spaceships
	case "glider":
		seeds = [][]int{
			[]int{1, 3},
			[]int{2, 1},
			[]int{2, 3},
			[]int{3, 2},
			[]int{3, 3},
		}
	case "lwss":
		seeds = [][]int{
			[]int{1, 2},
			[]int{1, 3},
			[]int{1, 4},
			[]int{1, 5},
			[]int{2, 1},
			[]int{2, 5},
			[]int{3, 5},
			[]int{4, 1},
			[]int{4, 4},
		}
	case "rand":
		rand.Seed(42)
		seeds = make([][]int, 0)
		for row := 0; row < Rows; row++ {
			for col := 0; col < Cols; col++ {
				if rand.Intn(10) > 7 {
					board[row][col] = true
				}
			}
		}

	}
	for _, seed := range seeds {
		board[seed[0]][seed[1]] = true
	}
}
