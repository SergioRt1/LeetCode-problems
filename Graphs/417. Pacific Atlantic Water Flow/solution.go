package main

import "fmt"

const (
	unseen = 0
	yes    = 1
	no     = 2
)

// dont work :S
func pacificAtlantic(heights [][]int) [][]int {
	var response [][]int
	pacific := make([][]int, len(heights))
	atlantic := make([][]int, len(heights))
	visit := make([][]bool, len(heights))

	for i := range heights {
		pacific[i] = make([]int, len(heights[i]))
		atlantic[i] = make([]int, len(heights[i]))
		visit[i] = make([]bool, len(heights[i]))
	}

	for i := range heights {
		for j := range heights[i] {
			if p, a := dfs(i, j, heights, pacific, atlantic, visit); p && a {
				response = append(response, []int{i, j})
			}
		}
	}

	return response
}

var (
	dx = []int{0, -1, 1, 0}
	dy = []int{-1, 0, 0, 1}
)

func dfs(i, j int, heights, pacific, atlantic [][]int, visit [][]bool) (p, a bool) {
	if pacific[i][j] != unseen && atlantic[i][j] != unseen {
		return pacific[i][j] == yes, atlantic[i][j] == yes
	}
	visit[i][j] = true
	for idx := range dx {
		nI := i + dx[idx]
		nJ := j + dy[idx]
		if nI < 0 || nJ < 0 {
			p = true
		} else if nI == len(heights) || nJ == len(heights[i]) {
			a = true
		} else if heights[i][j] >= heights[nI][nJ] && (!visit[nI][nJ] || (pacific[nI][nJ] != unseen && atlantic[nI][nJ] != unseen)) {
			sP, sA := dfs(nI, nJ, heights, pacific, atlantic, visit)
			p = p || sP
			a = a || sA
		}
	}
	if p {
		pacific[i][j] = yes
	} else {
		pacific[i][j] = no
	}
	if a {
		atlantic[i][j] = yes
	} else {
		atlantic[i][j] = no
	}
	for idx := range dx {
		nI := i + dx[idx]
		nJ := j + dy[idx]
		if 0 <= nI && nI != len(heights) && 0 <= nJ && nJ != len(heights[i]) &&
			heights[i][j] == heights[nI][nJ] {
			if p {
				pacific[nI][nJ] = yes
			}
			if a {
				atlantic[nI][nJ] = yes
			}
		}
	}

	return p, a
}

func main() {
	h := [][]int{
		{10, 10, 10},
		{10, 1, 10},
		{10, 10, 10},
	}
	fmt.Print(pacificAtlantic(h))
}
