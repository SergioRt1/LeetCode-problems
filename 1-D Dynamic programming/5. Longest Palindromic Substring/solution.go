package main

import "fmt"

// longestPalindrome Given a string s, return the longest palindromic substring in s.
func longestPalindrome(s string) string {
	// return recursiveHelper(s)
	return tabulation(s)
}

func tabulation(s string) string {
	memo := make([][]int, len(s))
	for i := range memo {
		memo[i] = make([]int, len(s))
	}
	// fill the main diagonal with 1
	for i := 0; i < len(s); i++ {
		memo[i][i] = 1
	}
	// fill the previews diagonal to the main with 0
	x := 1
	y := 0
	for ; x < len(s); x, y = x+1, y+1 {
		memo[x][y] = 0
	}

	globalBest := 1
	var bX, bY int

	// Decision best of either move i or j but don't take them, or take it if possible
	for d := len(s); d < len(s)*2-1; d++ {
		x = 0
		y = d - len(s) + 1

		for ; x < len(s) && y < len(s); x, y = x+1, y+1 {
			best := max(memo[x+1][y], memo[x][y-1]) // don't take it
			if s[x] == s[y] {
				mLengthIfTake := memo[x+1][y-1]
				if mLengthIfTake == y-x-1 {
					best = max(best, memo[x+1][y-1]+2) // take it
					if best > globalBest {
						globalBest = best
						bX = x
						bY = y
					}
				}
			}
			memo[x][y] = best

		}
	}
	return s[bX : bY+1]
}

func recursiveHelper(s string) string {
	memo := make([][]int, len(s))
	for i := range memo {
		memo[i] = make([]int, len(s))
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	best := dp(s, 0, len(s)-1, memo)
	if best == 1 {
		return s[:1]
	}
	for i := range memo {
		for j := range memo[i] {
			if memo[i][j] == best && best == j-i+1 {
				return s[i : j+1]
			}
		}
	}
	return s
}

// dp given a string s, return the max length of a palindrome inside s[i:j]
func dp(s string, i, j int, memo [][]int) int {
	// base cases
	if i == j {
		return 1
	}
	if i > j {
		return 0
	}
	if memo[i][j] != -1 {
		return memo[i][j]
	}
	// recursive case, recursion
	// Either move i or j but don't take them
	best := max(dp(s, i+1, j, memo), dp(s, i, j-1, memo))
	if s[i] == s[j] {
		// If they are equal, take values but make sure all elements between them are palindromes
		if mLengthIfTake := dp(s, i+1, j-1, memo); mLengthIfTake == j-i-1 {
			best = max(best, mLengthIfTake+2)
		}
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
	s := "ac"
	fmt.Println(longestPalindrome(s))
}
