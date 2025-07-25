package solution

// BEGIN

func FilterExpensiveOrders(orders []int, limit int) []int {
	result := make([]int, 0) // Инициализируем пустой срез
	for _, order := range orders {
		if order > limit {
			result = append(result, order)
		}
	}
	return result
}

// END
