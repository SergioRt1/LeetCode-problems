package main

import "fmt"

func isPalindrome(s string) bool {
	str := make([]int32, 0, len(s))
	for _, c := range s {
		if 'a' <= c && c <= 'z' {
			str = append(str, c-'a')
		} else if 'A' <= c && c <= 'Z' {
			str = append(str, c-'A')
		} else if '0' <= c && c <= '9' {
			str = append(str, c)
		}
	}
	l := len(str) - 1
	for i := range str {
		if str[i] != str[l-i] {
			return false
		}
	}
	return true
}

func main() {
	str := "0P0"
	fmt.Print(isPalindrome(str))
}
