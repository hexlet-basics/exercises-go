package solution

// BEGIN

// MaxSum gets sum of each nums slice and returns a slice with the max sum.
// If the slices are the same then the first one will be returned.
func MaxSum(nums1, nums2 []int) []int {
	if sum(nums1) >= sum(nums2) {
		return nums1
	}

	return nums2
}

func sum(nums []int) int {
	s := 0
	for _, n := range nums {
		s += n
	}

	return s
}

// END
