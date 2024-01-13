package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortByA(words []string) []string {
	sort.Slice(words, func(i, j int) bool {
		// fmt.Println(words)

		// Compare the number of "a"s in each word
		countA_i := strings.Count(words[i], "a")
		countA_j := strings.Count(words[j], "a")

		// fmt.Println(countA_i,countA_j)

		// fmt.Println(words[i], words[j])
		// fmt.Println(i, j)
		// fmt.Println(words)

		// fmt.Println(words[i] < words[j])

		// If the counts are equal, compare by word length

		// fmt.Println(len(words[i]) > len(words[j]))
		if countA_i == countA_j {
			return len(words[i]) > len(words[j])
		}

		// Otherwise, compare by the number of "a"s
		return countA_i > countA_j
	})

	// fmt.Println(words, "end")
	return words
}

func main() {
	words := []string{"banana", "apple", "cherry", "date", "apricot", "aaaaaaaa", "aaaa"}
	sortedWords := sortByA(words)

	fmt.Println(sortedWords, "end of")
}
