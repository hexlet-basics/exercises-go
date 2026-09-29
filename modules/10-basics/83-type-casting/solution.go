package solution

import (
	"fmt"
	"strconv"
)

// BEGIN

func BuildProfile(name string, age int, rating float64) string {
	ratingStr := strconv.FormatFloat(rating, 'f', 1, 64)
	return fmt.Sprintf("Name: %s, Age: %d, Rating: %s", name, age, ratingStr)
}

// END
