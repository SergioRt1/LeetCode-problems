package main

import (
	"fmt"
	"math"
)

// checkPalindrome checks if the given string is a palindrome
func checkPalindrome(s string) bool {
	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-i-1] {
			return false
		}
	}
	return true
}

// toBase converts the given number n to the specified base
func toBase(n, base int) string {
	var res string
	for n > 0 {
		res = string(n%base+'0') + res
		n /= base
	}
	return res
}

// findPalindromicBases finds all bases in which n is a palindrome
func findPalindromicBases(n int) []int {
	var bases []int

	// Special cases
	if n == 2 || n == 3 {
		return bases
	}

	// Check lower bases
	for base := 2; base <= int(math.Sqrt(float64(n))); base++ {
		if checkPalindrome(toBase(n, base)) {
			bases = append(bases, base)
		}
	}

	// Check higher bases
	for k := 1; k <= int(math.Sqrt(float64(n))); k++ {
		base := n - k
		if base > 1 && checkPalindrome(toBase(n, base)) {
			bases = append(bases, base)
		}
	}

	return bases
}

func main() {
	var n int
	fmt.Scan(&n) // read input

	bases := findPalindromicBases(n)
	if len(bases) == 0 {
		fmt.Println("*")
	} else {
		for _, base := range bases {
			fmt.Printf("%d ", base)
		}
		fmt.Println()
	}
}
