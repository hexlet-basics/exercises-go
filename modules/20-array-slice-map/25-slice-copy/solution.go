package solution

// BEGIN

// IntsCopy copies a slice of ints src with max length len.
// If maxLen is greater than src it returns a full copy of src.
// If maxLen is zero or negative it returns empty slice.
func IntsCopy(src []int, maxLen int) []int {
	if maxLen <= 0 {
		return []int{}
	}

	if maxLen > len(src) {
		maxLen = len(src)
	}

	cp := make([]int, maxLen)
	copy(cp, src)

	return cp[:maxLen]
}

// END
