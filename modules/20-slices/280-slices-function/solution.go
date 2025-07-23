package solution

// BEGIN

func AddDiscount(prices []int, discount int) []int {
	if discount <= 0 {
		return append([]int(nil), prices...) // без изменений, копируем
	}

	result := make([]int, len(prices))
	for i, price := range prices {
		result[i] = price - price*discount/100
		if result[i] < 0 {
			result[i] = 0
		}
	}
	return result
}

// END
