package solution

// BEGIN

import "slices"

func IntsCopy(src []int, maxLen int) []int {
	if maxLen <= 0 {
		return []int{}
	}

	if maxLen > len(src) {
		maxLen = len(src)
	}

	// Создаём копию только нужной части среза с помощью slices.Clone
	return slices.Clone(src[:maxLen])
}

// END
