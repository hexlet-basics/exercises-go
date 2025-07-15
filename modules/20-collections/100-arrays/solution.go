package solution

// BEGIN

// SafeWrite writes a value val in the array nums if the index i is in the array's boundaries.
func SafeWrite(nums [5]int, i, val int) [5]int {
	if i >= 0 && i < len(nums) {
		nums[i] = val
	}

	return nums
}

// END
