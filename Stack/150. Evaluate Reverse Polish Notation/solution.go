package main

import (
	"fmt"
	"strconv"
)

type operation = func(a, b int) int

var operations = map[uint8]operation{
	'*': func(a, b int) int { return a * b },
	'/': func(a, b int) int { return a / b },
	'+': func(a, b int) int { return a + b },
	'-': func(a, b int) int { return a - b },
}

func evalRPN(tokens []string) int {
	_, response := resolve(tokens)
	return response
}

// Recursive solution
func resolve(tokens []string) ([]string, int) {
	if len(tokens) > 0 {
		last := len(tokens) - 1
		newTokens, s := tokens[:last], tokens[last]

		switch op, isOperation := operations[s[0]]; {
		case isOperation && len(s) == 1:
			var a, b int
			newTokens, b = resolve(newTokens)
			newTokens, a = resolve(newTokens)
			return newTokens, op(a, b)
		default:
			value, _ := strconv.Atoi(s)
			return newTokens, value
		}
	}
	return nil, 0
}

func main() {
	s := []string{"3", "-4", "+"}
	fmt.Print(evalRPN(s))
}
