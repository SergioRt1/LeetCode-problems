package main

import (
	"fmt"
)

const (
	open  = 0
	close = 1
)

func longestValidParentheses(s string) int {
	i := 0
	for ; i <= len(s) && s[i] == ')'; i++ {
	}

	var dp func(idx, n, state, depth int) int
	dp = func(idx, n, state, depth int) int {
		if idx == len(s) {
			return 0
		}

		if s[idx] == '(' {
			return max(dp(idx+1, n+1, open, depth+1), dp(idx+1, 0, close, 0))
		} else { // is ')'
			if state == open {
				ns := open
				if n-1 == 0 {
					ns = close
				}
				return max(depth, dp(idx+1, n-1, ns, depth+1))
			} else {
				return 0
			}
		}
	}

	return dp(i, 0, close, 0)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s := ")()())"
	fmt.Println(longestValidParentheses(s))
}
