package main

import "fmt"

const dim = 9

func isValidSudoku(board [][]byte) bool {
	rows := make([]bool, 9)
	columns := make([]bool, 9)
	grid := make([]bool, 9)

	for i := 0; i < dim; i++ {
		clean(rows)
		clean(columns)
		clean(grid)
		for j := 0; j < dim; j++ {
			//Check rows
			if cell := board[i][j]; cell != '.' {
				num := cell - '1'
				if rows[num] {
					return false
				}
				rows[num] = true
			}
			//Check columns
			if cell := board[j][i]; cell != '.' {
				num := cell - '1'
				if columns[num] {
					return false
				}
				columns[num] = true
			}
			//Check grid
			ci := j/3 + (i/3)*3
			cj := j%3 + (i%3)*3
			if cell := board[ci][cj]; cell != '.' {
				num := cell - '1'
				if grid[num] {
					return false
				}
				grid[num] = true
			}
		}
	}

	return true
}

func clean(visit []bool) {
	for i := 0; i < dim; i++ {
		visit[i] = false
	}
}

func main() {
	board := [][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Print(isValidSudoku(board))
}
