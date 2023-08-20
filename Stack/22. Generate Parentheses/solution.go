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
		response = recursion(leftRemaining-1, leftAvailable+1, depth+1, stack)
	}
	if leftAvailable > 0 {
		stack[depth] = ')'
		if combinations := recursion(leftRemaining, leftAvailable-1, depth+1, stack); len(response) == 0 {
			response = combinations
		} else {
			response = append(response, combinations...)
		}

	}

	return response
}

func main() {
	n := 8
	fmt.Print(generateParenthesis(n))
}
