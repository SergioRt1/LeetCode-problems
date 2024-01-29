package main

import "fmt"

const (
	empty = iota
	fresh
	rotten
)

var (
	dx = []int{-1, 0, 0, 1}
	dy = []int{0, -1, 1, 0}
)

type Orange struct {
	x int
	y int
}

func orangesRotting(grid [][]int) int {
	queue := make([]Orange, 0)
	nodesInIteration := 0
	time := 0

	m, n := len(grid), len(grid[0])

	// find the rotten oranges
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == rotten {
				queue = append(queue, Orange{i, j})
				nodesInIteration++
			}
		}
	}
	// Run BFS
	for len(queue) > 0 {
		current := queue[0]
		nodesInIteration--
		queue = queue[1:] //dequeue

		// process
		for d := range dx {
			x, y := current.x+dx[d], current.y+dy[d]
			if 0 <= x && x < m && 0 <= y && y < n && grid[x][y] == fresh {
				grid[x][y] = rotten
				queue = append(queue, Orange{x, y})
			}
		}

		if nodesInIteration == 0 {
			time++
			nodesInIteration = len(queue)
		}
	}

	// validate grid
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == fresh {
				time = -1
			}
		}
	}
	if time >= 1 {
		time--
	}

	return time
}

func main() {
	grid := [][]int{
		{2, 0},
	}
	fmt.Println(orangesRotting(grid))
}
