package main

import "fmt"

type Node struct {
	x int
	y int
}
type nothing struct{}

var (
	dx = []int{-1, 0, 0, 1}
	dy = []int{0, -1, 1, 0}
)

func solve(board [][]byte) {
	m := len(board)
	n := len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				region := make(map[Node]nothing)
				current := Node{i, j}
				if dfs(current, region, board) {
					for node := range region {
						board[node.x][node.y] = 'X'
					}
				}
			}
		}
	}
}

func dfs(node Node, vis map[Node]nothing, board [][]byte) bool {
	vis[node] = nothing{}

	for d := range dx {
		x, y := node.x+dx[d], node.y+dy[d]
		newNode := Node{x, y}
		if isIn(newNode, len(board), len(board[0])) {
			_, visited := vis[newNode]
			if !visited && board[newNode.x][newNode.y] != 'X' {
				if !dfs(newNode, vis, board) {
					return false
				}
			}
		} else {
			return false
		}
	}

	return true
}

func isIn(node Node, m, n int) bool {
	return 0 <= node.x && node.x < m && 0 <= node.y && node.y < n
}

func printBoard(board [][]byte) {
	for i := range board {
		fmt.Println(string(board[i]))
	}
}

func main() {
	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	printBoard(board)
	solve(board)
	fmt.Println("###################")
	printBoard(board)
}
