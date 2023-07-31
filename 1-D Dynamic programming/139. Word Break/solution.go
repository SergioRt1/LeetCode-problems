package main

import "fmt"

const (
	isTrue        = 1
	isFalse       = -1
	notCalculated = 0

	maxLenWord = 21
)

func wordBreak(s string, wordDict []string) bool {
	memo := make([]int, len(s))
	wordMap := make(map[string]struct{}, len(wordDict))
	for _, word := range wordDict {
		wordMap[word] = struct{}{}
	}
	return dp(s, wordMap, 0, memo)
}

func dp(s string, wordDict map[string]struct{}, i int, memo []int) bool {
	if _, ok := wordDict[s[i:]]; ok {
		return true
	}
	if memo[i] != notCalculated {
		return memo[i] == isTrue
	}
	for j := i + 1; j < len(s) && j < i+maxLenWord; j++ {
		if _, ok := wordDict[s[i:j]]; ok {
			if dp(s, wordDict, j, memo) {
				memo[i] = isTrue
				return true
			}
		}
	}
	memo[i] = isFalse
	return false
}

func main() {
	s := "applepenapple"
	words := []string{"cat", "dog", "and", "cata"}
	fmt.Print(wordBreak(s, words))
}
