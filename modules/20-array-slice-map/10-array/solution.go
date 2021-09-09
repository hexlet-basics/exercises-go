package solution


// BEGIN

// SafeWrite writes a value val in the array arr if the index i is in the array's boundaries.
func SafeWrite(arr [5]int, i, val int) [5]int {
	if i < 0 || i >= len(arr) {
		return arr
	}

	arr[i] = val

	return arr
}

// END
