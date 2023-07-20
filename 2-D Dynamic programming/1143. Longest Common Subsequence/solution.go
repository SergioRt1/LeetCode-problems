package main

import "fmt"

func longestCommonSubsequence(text1 string, text2 string) int {
	return traversal(text1, text2)
	//return recursionHelper(text1, text2)
}

func traversal(text1, text2 string) int {
	memo := make([][]int, len(text1))
	for i := range memo {
		memo[i] = make([]int, len(text2))
	}

	for diag := 0; diag < len(text1)+len(text2); diag++ {
		var x, y int
		if diag < len(text1) {
			x = len(text1) - 1 - diag
			y = len(text2) - 1
		} else { // first value of diag == len(text1)
			x = 0
			y = len(text2) - 1 + len(text1) - diag
		}
		for ; x < len(text1) && y >= 0; x, y = x+1, y-1 {
			var best, take int

			existsDown := x+1 < len(text1)
			existsRight := y+1 < len(text2)
			if existsDown {
				best = max(best, memo[x+1][y])
			}
			if existsRight {
				best = max(best, memo[x][y+1])
			}
			if text1[x] == text2[y] {
				if existsDown && existsRight {
					take = memo[x+1][y+1]
				}
				best = max(best, take+1)
			}
			memo[x][y] = best
		}
	}
	return memo[0][0]
}

func recursionHelper(text1, text2 string) int {
	memo := make([][]int, len(text1))
	for i := range memo {
		memo[i] = make([]int, len(text2))
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	return recursion(text1, text2, 0, 0, memo)
}

func recursion(text1, text2 string, i, j int, memo [][]int) int {
	if i >= len(text1) || j >= len(text2) {
		return 0
	}
	if memo[i][j] != -1 {
		return memo[i][j]
	}

	// decision best of skipping i and j
	best := max(recursion(text1, text2, i+1, j, memo), recursion(text1, text2, i, j+1, memo))

	if text1[i] == text2[j] {
		best = max(best, recursion(text1, text2, i+1, j+1, memo)+1)
	}

	memo[i][j] = best
	return best
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Print(longestCommonSubsequence("c", "d"))
}
