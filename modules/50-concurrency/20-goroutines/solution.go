package solution

import (
	"time"
)

// BEGIN

// MaxSum gets sum of each nums slice and returns a slice with the max sum.
// If the slices are the same then the first one will be returned.
func MaxSum(nums1, nums2 []int) []int {
	s1, s2 := 0, 0

	go sumParallel(nums1, &s1)
	go sumParallel(nums2, &s2)

	time.Sleep(100 * time.Millisecond)

	if s1 >= s2 {
		return nums1
	}

	return nums2
}

func sumParallel(nums []int, res *int) {
	*res = sum(nums)
}

func sum(nums []int) int {
	s := 0
	for _, n := range nums {
		s += n
	}

	return s
}

// END
