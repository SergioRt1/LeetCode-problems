package main

import "fmt"

const delta = 'a' - 'A'

func isPalindrome(s string) bool {
	size := len(s)
	for i, j := 0, len(s)-1; i < j; {
		// skip non-alphanumeric characters
		for i < size && !isAlphanumeric(s[i]) {
			i++
		}
		for j >= 0 && !isAlphanumeric(s[j]) {
			j--
		}
		if i >= size || j < 0 {
			return true
		}

		if normalize(s[i]) != normalize(s[j]) {
			return false
		}

		//Advance
		i++
		j--
	}
	return true
}

func normalize(c uint8) uint8 {
	if 'A' <= c && c <= 'Z' {
		return c + delta
	}
	return c
}

func isAlphanumeric(c uint8) bool {
	if 'a' <= c && c <= 'z' {
		return true
	}
	if 'A' <= c && c <= 'Z' {
		return true
	}
	if '0' <= c && c <= '9' {
		return true
	}
	return false
}

func main() {
	s := ".,"
	fmt.Println(isPalindrome(s))
}
