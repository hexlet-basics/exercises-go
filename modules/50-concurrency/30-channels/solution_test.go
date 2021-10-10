package solution

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumWorker(t *testing.T) {
	a := assert.New(t)

	numsCh := make(chan []int)
	sumCh := make(chan int)

	go SumWorker(numsCh, sumCh)

	numsCh <- nil
	a.Equal(0, <-sumCh)

	numsCh <- []int{}
	a.Equal(0, <-sumCh)

	numsCh <- []int{10, 10, 10}
	a.Equal(30, <-sumCh)

	numsCh <- []int{500, 5, 10, 25}
	a.Equal(540, <-sumCh)
}
