package main

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
	//case 'pulsar':
	//seeds = [][]int{
	//for num1 in [2,7,9,14]
	//for num2 in [4,5,6,10,11,12]
	//seeds.push([num1, num2], [num2, num1])

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

	}
	for _, seed := range seeds {
		board[seed[0]][seed[1]] = true
	}
}

/*
    # Randomized
  case 'rand'
      seeds = []
      for row in [0...rows]
        for col in [0...cols]
          seeds.push([row, col]) if Math.random() > 0.7

  for rc in seeds
    [row, col] = rc
    board[row][col] = true

  return board
*/
