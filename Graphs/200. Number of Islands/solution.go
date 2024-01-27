package main

import "fmt"

var (
	dx = []int{-1, 0, 0, 1}
	dy = []int{0, -1, 1, 0}
)

func numIslands(grid [][]byte) int {
	count := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				count++
				dfs(i, j, grid)
			}
		}
	}

	return count
}

func dfs(i, j int, grid [][]byte) {
	grid[i][j] = '.'
	for d := range dx {
		x, y := i+dx[d], j+dy[d]
		if 0 <= x && x < len(grid) && 0 <= y && y < len(grid[0]) && grid[x][y] == '1' {
			dfs(x, y, grid)
		}
	}
}

type Node struct {
	i int
	j int
}

func bfs(n Node, grid [][]byte) {
	queue := []Node{n}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:] // dequeue
		grid[current.i][current.j] = '.'

		for d := range dx {
			x, y := current.i+dx[d], current.j+dy[d]
			if 0 <= x && x < len(grid) && 0 <= y && y < len(grid[0]) && grid[x][y] == '1' {
				queue = append(queue, Node{
					i: x,
					j: y,
				})
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
