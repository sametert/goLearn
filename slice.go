package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortByA(words []string) []string {
	sort.Slice(words, func(i, j int) bool {
		fmt.Println(words);

		// Compare the number of "a"s in each word
		countA_i := strings.Count(words[i], "a")
		countA_j := strings.Count(words[j], "a")

		// fmt.Println(countA_i,countA_j)

		// fmt.Println(words[i],words[j])
		// fmt.Println(i,j)
		// fmt.Println(words)

		//bool değer verir  true ise karşılaştırılan iki değer küçükten büyüğe doğru sıralanmıştır.
		// fmt.Println(words[i] < words[j])		

		// If the counts are equal, compare by word length
		//kelimenin içerisinde ki a harfi eşitse

		// fmt.Println(len(words[i]) > len(words[j]))
		if countA_i == countA_j {
			return len(words[i]) > len(words[j])
		}

		// Otherwise, compare by the number of "a"s
		return countA_i > countA_j
	})

	fmt.Println(words,"end");
	return words
}

func main() {
	words := []string{"banana", "apple", "cherry", "date", "apricot"}
	sortedWords := sortByA(words)

	fmt.Println(sortedWords,"end of")
}

// func main() {
// 	numbers := []int{5, 2, 9, 1, 5, 6}
	
// 	// Kucukten buyuge siralama
// 	sort.Slice(numbers, func(i, j int) bool {
// 		return numbers[i] < numbers[j]
// 	})

	
// 	fmt.Println("Küçükten Büyüğe Sıralı Dilim:", numbers)

// 	// Buyukten kucuge siralama
// 	sort.Slice(numbers, func(i, j int) bool {
// 		return numbers[i] > numbers[j]
// 	})
// 	fmt.Println("Büyükten Küçüğe Sıralı Dilim:", numbers)
// }