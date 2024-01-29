package main

import "fmt"

func findOrder(numCourses int, prerequisites [][]int) []int {
	visited := make([]bool, numCourses)
	isParent := make([]bool, numCourses)
	order := make([]int, 0, numCourses)
	g := make([][]int, numCourses) // graph represented as Adj list

	for _, pre := range prerequisites {
		g[pre[1]] = append(g[pre[1]], pre[0])
	}

	var ok bool
	for c := 0; c < numCourses; c++ {
		if !visited[c] {
			if order, ok = dfs(c, g, visited, isParent, order); !ok {
				return nil
			}
		}
	}

	for i, j := 0, numCourses-1; i < j; i, j = i+1, j-1 {
		order[i], order[j] = order[j], order[i]
	}

	return order
}

func dfs(course int, g [][]int, visited []bool, isParent []bool, order []int) ([]int, bool) {
	visited[course] = true
	isParent[course] = true
	for _, c := range g[course] {
		if !visited[c] {
			var ok bool
			// order
			order, ok = dfs(c, g, visited, isParent, order)
			if !ok {
				return nil, false
			}
		} else if isParent[c] { // cycle detected
			return nil, false
		}
	}

	isParent[course] = false

	return append(order, course), true
}

func main() {
	num := 6
	pre := [][]int{
		{2, 4}, {1, 0}, {1, 5}, {3, 1}, {3, 2},
	}
	fmt.Println(findOrder(num, pre))
}
