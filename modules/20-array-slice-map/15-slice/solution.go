package solution

// BEGIN

func Reverse(slice []int) []int {
	reversed := make([]int, len(slice))

	for i := len(slice) - 1; i >= 0; i-- {
		reversed[len(slice)-1-i] = slice[i]
	}

	return reversed
}

// END
