package main

import (
	"fmt"
)

// BEGIN

// Структура Account
type Account struct {
	Owner   string
	Balance int
}

// Функция пополнения счёта
func Deposit(acc *Account, amount int) {
	if amount > 0 {
		acc.Balance += amount
	}
}

func main() {
	acc := Account{Owner: "Alice", Balance: 100}
	Deposit(&acc, 50)
	fmt.Println(acc.Balance) // 150
}

// END
