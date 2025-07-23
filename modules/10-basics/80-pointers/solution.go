package solution

// BEGIN

type Account struct {
	Owner   string
	Balance int
}

func Deposit(acc *Account, amount int) {
	if amount > 0 {
		acc.Balance += amount
	}
}

// END
