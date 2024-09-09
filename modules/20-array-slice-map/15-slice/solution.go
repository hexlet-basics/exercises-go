package solution

// BEGIN

func RemoveFirstElement(slice []int) []int {
	if len(slice) == 0 {
		return slice
	}

	return slice[1:]
}

// END
