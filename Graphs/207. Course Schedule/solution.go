package main

import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {
	visited := make([]bool, numCourses)
	parent := make([]bool, numCourses)
	graph := make([][]int, numCourses)

	for _, pre := range prerequisites {
		graph[pre[1]] = append(graph[pre[1]], pre[0])
	}

	for i := 0; i < numCourses; i++ {
		if !visited[i] {
			if !dfs(i, graph, visited, parent) {
				return false
			}
		}
	}

	return true
}

func dfs(course int, graph [][]int, visited []bool, parent []bool) bool {
	parent[course] = true
	visited[course] = true
	for _, c := range graph[course] {
		if !visited[c] {
			if !dfs(c, graph, visited, parent) {
				return false
			}
		} else if parent[c] {
			return false
		}
	}

	parent[course] = false
	return true
}

func main() {
	num := 5
	pre := [][]int{
		{1, 0}, {0, 1},
	}
	fmt.Println(canFinish(num, pre))
}
