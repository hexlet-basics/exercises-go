package solution

import "strconv"

// BEGIN

func BuildProfile(name string, age int, rating float64) string {
	ageStr := strconv.Itoa(age)
	ratingStr := fmt.Sprintf("%.1f", rating)
	return "Name: " + name + ", Age: " + ageStr + ", Rating: " + ratingStr
}

// END
