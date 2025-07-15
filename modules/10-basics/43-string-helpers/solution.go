package solution

import (
	"strings"
)

// BEGIN

func GetHiddenCard(cardNumber string, starsCount int) string {
	return strings.Repeat("*", starsCount) + cardNumber[len(cardNumber)-4:]
}

// END
