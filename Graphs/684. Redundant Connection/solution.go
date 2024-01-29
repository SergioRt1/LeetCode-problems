package main

import "fmt"

type nothing struct{}

func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	visited := make([]bool, n+1)
	parent := make([]int, n)
	g := make([][]int, n+1)

	// build graph
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
	}

	//find cycle
	start, depth := findCycle(1, 0, g, parent, visited)
	for i := 0; i < len(parent); i++ {
		if parent[i] == start {
			parent = parent[i:depth]
			break
		}
	}

	cycleG := make([]map[int]nothing, n+1)
	cycleLen := len(parent)
	for i := 0; i < cycleLen; i++ {
		node1 := parent[i]
		node2 := parent[(i+1)%cycleLen]
		if cycleG[node1] == nil {
			cycleG[node1] = make(map[int]nothing)
		}
		if cycleG[node2] == nil {
			cycleG[node2] = make(map[int]nothing)
		}

		cycleG[node1][node2] = nothing{}
		cycleG[node2][node1] = nothing{}
	}

	var response []int

	for _, e := range edges {
		if _, isInCycle := cycleG[e[0]][e[1]]; isInCycle {
			response = e
		}
	}

	return response
}

func findCycle(node, depth int, g [][]int, parent []int, visited []bool) (int, int) {
	visited[node] = true
	parent[depth] = node
	for _, n := range g[node] {
		if depth <= 0 || parent[depth-1] != n {
			if !visited[n] {
				cycle, d := findCycle(n, depth+1, g, parent, visited)
				if cycle != 0 {
					return cycle, d
				}
			} else {
				return n, depth + 1
			}
		}
	}

	return 0, 0
}

func main() {
	e := [][]int{{1, 2}, {1, 3}, {2, 3}}
	//e := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}

	fmt.Println(findRedundantConnection(e))
}
