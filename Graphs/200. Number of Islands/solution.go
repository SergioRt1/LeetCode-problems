package main

import "fmt"

// numIslands Use a dfs traversal to mark an entire island of '1' as visited '.', checks all cells in the grid counting the number of islands
func numIslands(grid [][]byte) int {
	nIslands := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				nIslands++
				dfs(grid, i, j)
			}
		}
	}
	return nIslands
}

var dirI = []int{-1, 0, 0, 1}
var dirJ = []int{0, -1, 1, 0}

func dfs(grid [][]byte, i, j int) {
	if grid[i][j] == '1' {
		grid[i][j] = '.'
		for idx := range dirI {
			nI := i + dirI[idx]
			nJ := j + dirJ[idx]
			if 0 <= nI && nI < len(grid) && 0 <= nJ && nJ < len(grid[i]) {
				dfs(grid, nI, nJ)
			}
		}
	}
}

func main() {
	grid := [][]byte{
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	fmt.Print(numIslands(grid))
}
