package main

import "fmt"

// counts the number of occurrences of a letter in the word for each string and then compares them
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	lettersS := countLetters(s)
	lettersT := countLetters(t)
	for i := range lettersT {
		if lettersT[i] != lettersS[i] {
			return false
		}
	}
	return true
}

func countLetters(s string) []int {
	alf := make([]int, 26)
	for _, l := range s {
		alf[l-'a']++
	}
	return alf
}

func main() {
	s := "nl"
	t := "cx"
	fmt.Print(isAnagram(s, t))
}
