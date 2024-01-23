package solution

// BEGIN (write your solution here)
// MostPopularWord returns most popular word from the words slice.
// If there are multiple popular words it returns the first one depending on the words slice order.
func MostPopularWord(words []string) string {
	wordsCount := make(map[string]int, 0)
	mostPopWord := ""
	max := 0

	for _, word := range words {
		wordsCount[word]++
		if wordsCount[word] > max {
			max = wordsCount[word]
		}
	}

	for _, word := range words {
		if wordsCount[word] == max {
			mostPopWord = word
			break
		}
	}
	
	return mostPopWord
}

// END
