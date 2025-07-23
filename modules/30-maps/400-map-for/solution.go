package solution

import "strings"

// BEGIN

func CountWords(text string) map[string]int {
	result := make(map[string]int)
	words := strings.FieldsSeq(strings.ToLower(text))

	for word := range words {
		// Удаляем знаки пунктуации с конца и начала
		word = strings.Trim(word, ".,!?;:")
		if word != "" {
			result[word]++
		}
	}
	return result
}

// END
