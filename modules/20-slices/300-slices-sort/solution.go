package solution

import "slices"

// BEGIN

type Order struct {
	CustomerID int
	Price      int
}

func SortOrdersByCustomerID(orders []Order) []Order {
	result := slices.Clone(orders)
	slices.SortFunc(result, func(a, b Order) int {
		if a.CustomerID == b.CustomerID {
			return a.Price - b.Price
		}
		return a.CustomerID - b.CustomerID
	})
	return result
}

// END
