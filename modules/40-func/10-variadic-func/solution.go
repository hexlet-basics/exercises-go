package solution

// BEGIN

// MergeNumberLists merges all the number lists into the single one and returns it.
// It preserves the order of the number list items.
func MergeNumberLists(numberLists ...[]int) []int {
	mergedCap := 0
	for i := 0; i < len(numberLists); i++ {
		mergedCap += len(numberLists[i])
	}

	merged := make([]int, 0, mergedCap)
	for _, nl := range numberLists {
		merged = append(merged, nl...)
	}

	return merged
}

// END
