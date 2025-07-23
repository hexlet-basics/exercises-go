package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterExpensiveOrders(t *testing.T) {
	a := assert.New(t)
	a.Equal([]int{120, 400, 220}, FilterExpensiveOrders([]int{120, 35, 70, 400, 15, 220, 90}, 100))
	a.Equal([]int{}, FilterExpensiveOrders([]int{}, 100))
	a.Equal([]int{}, FilterExpensiveOrders([]int{10, 20, 30}, 100))
	a.Equal([]int{150, 200}, FilterExpensiveOrders([]int{150, 200}, 100))
	a.Equal([]int{}, FilterExpensiveOrders([]int{100, 100}, 100)) // ровно limit не включается
}
