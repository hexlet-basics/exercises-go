package solution

// BEGIN

// Структура Product
type Product struct {
	Name  string
	Price int
}

// NewDiscountedProduct создает новый товар с учетом скидки.
func NewDiscountedProduct(name string, price int, discount int) *Product {
	if discount < 0 {
		discount = 0
	}
	if discount > 100 {
		discount = 100
	}
	finalPrice := price - (price * discount / 100)
	return &Product{Name: name, Price: finalPrice}
}

// END
