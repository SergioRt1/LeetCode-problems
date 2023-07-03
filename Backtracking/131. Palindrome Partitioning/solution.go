package main

import (
	"fmt"
)

func partition(s string) [][]string {
	return helper(s, 0)
}

// helper also pass a currentIdx which is the current deep of the recursion, this helps build the response array with the final size
func helper(s string, currentIdx int) [][]string {
	response := make([][]string, 0, len(s))
	// base case
	if len(s) == 0 {
		return append(response, make([]string, currentIdx))
	}
	// recursive case
	// decision: find all possible prefix-palindromes in s, then find partitions for the 'rest' of s
	for _, palindrome := range getPrefixPalindromes(s) {
		for _, subPartition := range helper(s[len(palindrome):], currentIdx+1) {
			subPartition[currentIdx] = palindrome
			response = append(response, subPartition)
		}
	}
	return response
}

func getPrefixPalindromes(s string) []string {
	response := make([]string, 0)
	for i := 0; i < len(s); i++ {
		isPalindrome := true
		for j, k := 0, i; j <= i>>1 && isPalindrome; j, k = j+1, k-1 {
			isPalindrome = isPalindrome && s[j] == s[k]
		}
		if isPalindrome {
			response = append(response, s[:i+1])
		}
	}
	return response
}

func main() {
	s := "aab"
	fmt.Print(partition(s))
}
