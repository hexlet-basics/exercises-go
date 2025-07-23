package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeposit(t *testing.T) {
	tests := []struct {
		name     string
		initial  int
		amount   int
		expected int
	}{
		{"Positive deposit", 100, 50, 150},
		{"Zero deposit", 100, 0, 100},
		{"Negative deposit", 100, -20, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			acc := Account{Owner: "Test", Balance: tt.initial}
			Deposit(&acc, tt.amount)
			assert.Equal(t, tt.expected, acc.Balance, "Balance mismatch in test case: %s", tt.name)
		})
	}
}
