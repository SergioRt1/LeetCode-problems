package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	lenS, ans := len(s), 0
	if lenS <= 1 {
		return lenS
	}
	i, j := 0, 1
	set := make(map[byte]struct{})
	set[s[i]] = struct{}{}
	for i < lenS && j < lenS {
		if _, isIn := set[s[j]]; isIn {
			delete(set, s[i])
			i++
		} else {
			set[s[j]] = struct{}{}
			j++
		}

		if j-i > ans {
			ans = j - i
		}
	}

	return ans
}

func main() {
	s := "abcabcbb"
	fmt.Print(lengthOfLongestSubstring(s))
}
