package main

import "fmt"

func exist(board [][]byte, word string) bool {
	for i := range board {
		for j := range board[i] {
			if dfs(board, word, i, j) {
				return true
			}
		}
	}
	return false
}

var dI = []int{-1, 0, 0, 1}
var dJ = []int{0, -1, 1, 0}

func dfs(board [][]byte, word string, i, j int) bool {
	if c := board[i][j]; word[0] == c {
		if len(word) == 1 {
			return true
		}
		board[i][j] = '.'
		//check surrounding cells
		for k := range dI {
			nI := i + dI[k]
			nJ := j + dJ[k]
			if 0 <= nI && nI < len(board) && 0 <= nJ && nJ < len(board[i]) && dfs(board, word[1:], nI, nJ) {
				return true
			}
		}
		board[i][j] = c
	}
	return false
}

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCB"
	fmt.Print(exist(board, word))
}
