package solution

// BEGIN

func FilterExpensiveOrders(orders []int, limit int) []int {
	var result []int
	for _, order := range orders {
		if order > limit {
			result = append(result, order)
		}
	}
	return result
}

// END
