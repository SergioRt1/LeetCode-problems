package main

import "fmt"

func uniquePaths(m int, n int) int {
	//return recursiveHelper(m, n)
	return tabulation(m, n)
}

func tabulation(m, n int) int {
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
	}
	memo[0][0] = 1

	// Traversal over the diagonals
	for diag := 1; diag < m+n-1; diag++ {
		var x, y int
		if diag < n {
			x = 0
			y = diag
		} else { // first time diag == n
			x = diag - n + 1
			y = n - 1
		}
		for ; x < m && y >= 0; x, y = x+1, y-1 {
			var left, up int
			if x-1 >= 0 {
				up = memo[x-1][y]
			}
			if y-1 >= 0 {
				left = memo[x][y-1]
			}
			memo[x][y] = left + up
		}
	}

	return memo[m-1][n-1]
}

func recursiveHelper(m int, n int) int {
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	return recursiveDP(m-1, n-1, memo)
}

func recursiveDP(m, n int, memo [][]int) int {
	if m == 0 && n == 0 {
		return 1
	}
	if m < 0 || n < 0 {
		return 0
	}
	if memo[m][n] != -1 {
		return memo[m][n]
	}
	memo[m][n] = recursiveDP(m, n-1, memo) + recursiveDP(m-1, n, memo)
	return memo[m][n]
}

func main() {
	fmt.Print(uniquePaths(3, 7))
}
