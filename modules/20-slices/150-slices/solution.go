package solution

// BEGIN

func AddDiscount(prices []float64, discount float64) []float64 {
	if len(prices) == 0 {
		return prices
	}

	prices[len(prices)-1] = prices[len(prices)-1] * (1 - discount/100)

	return prices
}

// END
