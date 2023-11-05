package main

import "fmt"

func main() {
	var n, g int

	fmt.Scan(&n)
	words := make([]string, n)
	hasLetters := make([][5]bool, n)
	hasWord := make([]bool, 26)
	for i := 0; i < n; i++ {
		fmt.Scan(&words[i])
		if i == 0 {
			for j, c := range words[0] {
				hasWord[c-'a'] = true
				hasLetters[0][j] = true
			}
		} else {
			for j, c := range words[i] {
				hasLetters[i][j] = hasWord[c-'a']
			}
		}
	}

	fmt.Scan(&g)
	guesses := make([]string, g)
	for i := 0; i < g; i++ {
		fmt.Scan(&guesses[i])
	}
	answer := words[0]

	for i := range guesses {
		count := 0
	wordsLoop:
		for j, word := range words {

			for k, c := range guesses[i] {
				switch c {
				case '*':
					if answer[k] != word[k] {
						continue wordsLoop
					}
				case 'X':
					if answer[k] == word[k] || hasLetters[j][k] {
						continue wordsLoop
					}
				case '!':
					if answer[k] == word[k] || !hasLetters[j][k] {
						continue wordsLoop
					}

				}
			}
			count++
		}
		fmt.Println(count)
	}
}
