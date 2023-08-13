package main

import (
	"fmt"
)

func generateParenthesis(n int) []string {
	stack := make([]byte, n*2)

	return recursion(n, 0, 0, stack)
}

func recursion(leftRemaining, leftAvailable, depth int, stack []byte) []string {
	if leftRemaining == 0 && leftAvailable == 0 {
		return []string{string(stack)}
	}
	var response []string

	if leftRemaining > 0 {
		stack[depth] = '('
		response = append(response, recursion(leftRemaining-1, leftAvailable+1, depth+1, stack)...)
	}
	if leftAvailable > 0 {
		stack[depth] = ')'
		response = append(response, recursion(leftRemaining, leftAvailable-1, depth+1, stack)...)
	}

	return response
}

func main() {
	n := 8
	fmt.Print(generateParenthesis(n))
}
