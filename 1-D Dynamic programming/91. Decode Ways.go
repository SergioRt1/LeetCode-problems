package main

import "fmt"

func numDecodings(s string) int {
	return tabulation(s)
}

func tabulation(s string) int {
	last := len(s) - 1
	memo := make([]int, len(s)+2)
	memo[last+1] = 1
	memo[last+2] = 1

	for i := last; i >= 0; i-- {
		if s[i] > '0' {
			memo[i] = memo[i+1]
		}
		if i+1 < len(s) {
			if s[i] == '1' || (s[i] == '2' && s[i+1] <= '6') {
				memo[i] += memo[i+2]
			}
		}
	}

	return memo[0]
}

func recursionHelper(s string) int {
	memo := make([]int, len(s))
	for i := range memo {
		memo[i] = -1
	}
	return dp(s, 0, memo)
}

func dp(s string, i int, memo []int) int {
	if i >= len(s) {
		return 1
	}
	if memo[i] != -1 {
		return memo[i]
	}
	var resp int
	if s[i] > '0' {
		resp = dp(s, i+1, memo)
	}
	if i+1 < len(s) {
		if s[i] == '1' || (s[i] == '2' && s[i+1] <= '6') {
			resp += dp(s, i+2, memo)
		}
	}
	memo[i] = resp
	return resp
}

func main() {
	s := "111111111111111111111111111111111111111111111"
	fmt.Print(numDecodings(s))
}
