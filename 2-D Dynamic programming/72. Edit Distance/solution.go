package main

import "fmt"

const (
	Unseen = -1
)

func minDistance(word1, word2 string) int {
	memo := make([][]int, len(word1))
	for i := range memo {
		memo[i] = make([]int, len(word2))
		for j := range memo[i] {
			memo[i][j] = Unseen
		}
	}

	return dp(word1, word2, 0, 0, memo)
}

func dp(word1, word2 string, i, j int, memo [][]int) int {
	// base cases
	if i == len(word1) && j == len(word2) { // done Ok
		return 0
	}
	if i == len(word1) {
		return len(word2) - j
	}
	if j == len(word2) {
		return len(word1) - i
	}
	if memo[i][j] != Unseen {
		return memo[i][j]
	}
	// recursive case
	replace := dp(word1, word2, i+1, j+1, memo)
	if word1[i] != word2[j] {
		replace += 1
	}
	deleteOrInsert := min(dp(word1, word2, i+1, j, memo), dp(word1, word2, i, j+1, memo)) + 1
	memo[i][j] = min(replace, deleteOrInsert)

	return memo[i][j]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	w1 := "horse"
	w2 := "ros"
	fmt.Println(minDistance(w1, w2))
}
