package main

import "fmt"

func isValid(s string) bool {
	brackets := make([]uint8, 0)
	validPairs := map[uint8]uint8{
		')': '(',
		'}': '{',
		']': '[',
	}

	for i := range s {
		switch c := s[i]; c {
		case '(', '{', '[':
			brackets = append(brackets, c)
		case ')', '}', ']':
			if len(brackets) == 0 || brackets[len(brackets)-1] != validPairs[c] {
				return false
			} else {
				brackets = brackets[:len(brackets)-1]
			}
		}
	}

	return len(brackets) == 0
}

func main() {
	s := "{"
	fmt.Print(isValid(s))
}
