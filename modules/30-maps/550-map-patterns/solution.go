package solution

// BEGIN

func CountLanguages(users map[string]string) map[string]int {
	result := make(map[string]int)
	for _, lang := range users {
		result[lang]++
	}
	return result
}

// END
