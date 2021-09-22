package solution

// BEGIN

type popularityIndex struct {
	firstIndex int
	popularity int
}

// MostPopularWord returns most popular word from the words slice.
// If there are multiple popular words it returns the first one depending on the words slice order.
func MostPopularWord(words []string) string {
	wordsPopularity := make(map[string]popularityIndex)

	for i, w := range words {
		pi, ok := wordsPopularity[w]
		if !ok {
			wordsPopularity[w] = popularityIndex{
				firstIndex: i,
				popularity: 1,
			}
			continue
		}

		pi.popularity++

		wordsPopularity[w] = pi
	}

	var (
		mostPopWord     string
		firstPopWordIdx int
		maxPop          int
	)

	for w, pi := range wordsPopularity {
		if pi.popularity > maxPop {
			maxPop = pi.popularity
			firstPopWordIdx = pi.firstIndex
			mostPopWord = w
			continue
		}

		if pi.popularity == maxPop && pi.firstIndex < firstPopWordIdx {
			maxPop = pi.popularity
			firstPopWordIdx = pi.firstIndex
			mostPopWord = w
		}
	}

	return mostPopWord
}

// END
