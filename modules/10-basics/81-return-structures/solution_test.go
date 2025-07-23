package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDiscountedProduct(t *testing.T) {
	tests := []struct {
		name      string
		price     int
		discount  int
		wantPrice int
	}{
		{"No discount", 1000, 0, 1000},
		{"10% discount", 1000, 10, 900},
		{"50% discount", 1000, 50, 500},
		{"100% discount", 1000, 100, 0},
		{"Negative discount", 1000, -10, 1000},
		{"Over 100% discount", 1000, 150, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewDiscountedProduct("Test", tt.price, tt.discount)
			assert.Equal(t, tt.wantPrice, p.Price, "Price for test case '%s' did not match", tt.name)
		})
	}
}
