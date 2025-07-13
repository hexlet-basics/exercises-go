package solution

// BEGIN

func RepeatWord(word string, times int) string {
	result := ""
	for range times {
		result += word
	}
	return result
}

// END
