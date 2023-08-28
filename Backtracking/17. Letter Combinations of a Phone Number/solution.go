package main

import (
	"fmt"
	"strings"
)

var phone = map[byte][]byte{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	res := helper(digits)
	response := make([]string, 0, len(res))
	for i := range res {
		response = append(response, reverse(res[i].String()))
	}
	return response
}

func helper(digits string) []strings.Builder {
	response := make([]strings.Builder, 0)
	if len(digits) == 0 {
		return response
	}

	// recursive case / decision
	combinations := helper(digits[1:])
	if len(combinations) == 0 {
		combinations = append(combinations, strings.Builder{})
	}
	for i := range combinations {
		for _, letter := range phone[digits[0]] {
			newCombination := strings.Builder{}
			newCombination.WriteString(combinations[i].String())

			newCombination.WriteByte(letter)
			response = append(response, newCombination)
		}
	}
	return response
}

func reverse(s string) string {
	r := []byte(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return string(r)
}

func main() {
	digits := "23"
	fmt.Print(letterCombinations(digits))
}
