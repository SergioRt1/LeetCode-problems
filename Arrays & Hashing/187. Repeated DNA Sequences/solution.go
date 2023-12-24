package main

import (
	"fmt"
)

const mask = 1<<20 - 1

func findRepeatedDnaSequences(s string) []string {
	code := 0
	set := make(map[int]int)
	for i := range s {
		code += mapping(s[i])
		if i >= 9 {
			set[code] += 1
		}

		code <<= 2
		code &= mask
	}
	response := make([]string, 0)
	for c, count := range set {
		if count > 1 {
			response = append(response, decode(c))
		}
	}

	return response
}

func mapping(l uint8) int {
	switch l {
	case 'A':
		return 0
	case 'C':
		return 1
	case 'G':
		return 2
	case 'T':
		return 3
	}
	return 0
}

func decode(code int) string {
	sequence := make([]byte, 10)
	i := 9
	for code > 0 {
		sequence[i] = decodeMapping(code & 0b11)
		code >>= 2
		i--
	}
	for i >= 0 {
		sequence[i] = 'A'
		i--
	}

	return string(sequence)
}

func decodeMapping(c int) byte {
	switch c {
	case 0:
		return 'A'
	case 1:
		return 'C'
	case 2:
		return 'G'
	case 3:
		return 'T'
	}
	return 'A'
}

func main() {
	s := "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
	fmt.Println(findRepeatedDnaSequences(s))
}
