package solution

// BEGIN

func RepeatWord(word string, times int) string {
	result := ""
	for i := 0; i < times; i++ {
		result += word
	}
	return result
}

// END
