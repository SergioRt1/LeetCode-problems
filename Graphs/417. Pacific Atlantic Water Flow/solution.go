package main

import "fmt"

var (
	dx = []int{0, -1, 1, 0}
	dy = []int{-1, 0, 0, 1}
)

type node struct {
	x int
	y int
}

type nothing struct{}

func pacificAtlantic(heights [][]int) [][]int {
	pacific := map[node]nothing{}
	atlantic := map[node]nothing{}
	response := make([][]int, 0)
	m := len(heights)
	n := len(heights[0])
	for i := 0; i < m; i++ {
		dfs(i, 0, heights, pacific)    //left
		dfs(i, n-1, heights, atlantic) //right
	}
	for j := 0; j < n; j++ {
		dfs(0, j, heights, pacific)    //top
		dfs(m-1, j, heights, atlantic) //bottom
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			current := node{i, j}
			_, pacificOk := pacific[current]
			_, atlanticOk := atlantic[current]
			if pacificOk && atlanticOk {
				response = append(response, []int{i, j})
			}
		}
	}

	return response
}

func dfs(i, j int, heights [][]int, visited map[node]nothing) {
	currentHeight := heights[i][j]
	current := node{i, j}
	visited[current] = nothing{}
	for d := range dx {
		x, y := i+dx[d], j+dy[d]
		if 0 <= x && x < len(heights) && 0 <= y && y < len(heights[0]) {
			neig := node{x, y}
			_, vis := visited[neig]
			if !vis && currentHeight <= heights[x][y] {
				dfs(x, y, heights, visited)
			}
		}
	}
}

func main() {
	h := [][]int{
		{1, 2, 2, 3, 5},
		{3, 2, 3, 4, 4},
		{2, 4, 5, 3, 1},
		{6, 7, 1, 4, 5},
		{5, 1, 1, 2, 4},
	}
	fmt.Print(pacificAtlantic(h))
}
