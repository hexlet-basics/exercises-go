package solution

// BEGIN

// Map iterates through the strs slice and modifies each element via mapFunc. The func is safe and strs won't be modified.
func Map(strs []string, mapFunc func(s string) string) []string {
	mapped := make([]string, len(strs))
	for i, s := range strs {
		mapped[i] = mapFunc(s)
	}

	return mapped
}

// END
