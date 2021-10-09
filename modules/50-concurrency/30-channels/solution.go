package solution

// BEGIN

// SumWorker runs a worker that calculates a sum of nums from the numsCh and returns a result in the sumCh.
func SumWorker(numsCh chan []int, sumCh chan int) {
	for nums := range numsCh {
		sumCh <- sum(nums)
	}
}

func sum(nums []int) int {
	s := 0
	for _, n := range nums {
		s += n
	}

	return s
}

// END
