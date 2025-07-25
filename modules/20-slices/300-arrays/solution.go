package solution

// BEGIN

func SafeWrite(nums [5]int, i, val int) [5]int {
	if i >= 0 && i < len(nums) {
		nums[i] = val
	}

	return nums
}

// END
