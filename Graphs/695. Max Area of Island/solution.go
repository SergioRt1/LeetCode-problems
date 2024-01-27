package main

import "fmt"

var (
	dx = []int{-1, 0, 0, 1}
	dy = []int{0, -1, 1, 0}
)

func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				maxArea = max(maxArea, dfs(i, j, grid))
			}
		}
	}

	return maxArea
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dfs(i, j int, grid [][]int) int {
	grid[i][j] = -1
	nodes := 1
	for d := range dx {
		x, y := i+dx[d], j+dy[d]
		if 0 <= x && x < len(grid) && 0 <= y && y < len(grid[0]) && grid[x][y] == 1 {
			nodes += dfs(x, y, grid)
		}
	}
	return nodes
}

func main() {
	grid := [][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}
	fmt.Println(maxAreaOfIsland(grid))
	for i := range grid {
		fmt.Println(grid[i])
	}
}
