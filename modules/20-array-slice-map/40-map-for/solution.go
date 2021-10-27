package solution

// BEGIN
func MostPopularWord(words []string) string {
	wordsCount := make(map[string]int, 0)
	result := ""
	max := 0

	for _, word := range words {
		wordsCount[word]++
		if wordsCount[word] > max {
			max = wordsCount[word]
			result = word
		}
	}
	return result
}
// END
