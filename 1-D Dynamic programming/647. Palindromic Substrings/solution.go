package main

import "fmt"

func countSubstrings(s string) int {
	mPal := make([][]int, len(s))

	for i := range mPal {
		mPal[i] = make([]int, len(s))
	}

	// isPalindrome(s, 0, len(s)-1, mPal)

	return tabulationIsPal(s, mPal)
}

// tabulationIsPal starts bottom left and move over the diagonals to the top right corner, return the count
func tabulationIsPal(s string, mPal [][]int) int {
	count := 0
	for d := len(s) - 2; d < len(s)*2-1; d++ {
		var x, y int
		if d < len(s) {
			x = len(s) - d - 1
			y = 0
		} else {
			x = 0
			y = d - len(s) + 1
		}
		for ; y < len(s) && x < len(s); x, y = x+1, y+1 {
			if x > y { // base case
				mPal[x][y] = 1
			} else if x == y || (s[x] == s[y] && mPal[x+1][y-1] == 1) { // single letter or is pal
				mPal[x][y] = 1
				count++
			}
		}
	}
	return count
}

// isPalindrome uses DP to calculate al possible palindromes in s[i:j]
// returns 1 if is palindrome and -1 if not
func isPalindrome(s string, i, j int, mPal [][]int) int {
	if i >= j {
		return 1
	}
	if mPal[i][j] != 0 { // a value of 0 means that it has not yet been calculated
		return mPal[i][j]
	}
	if s[i] == s[j] && isPalindrome(s, i+1, j-1, mPal) == 1 {
		mPal[i][j] = 1
	} else {
		mPal[i][j] = -1
	}
	// possible substrings
	isPalindrome(s, i+1, j, mPal)
	isPalindrome(s, i, j-1, mPal)

	return mPal[i][j]
}

func main() {
	s := "aaa"
	fmt.Print(countSubstrings(s))
}
