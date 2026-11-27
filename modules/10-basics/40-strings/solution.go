package solution

// BEGIN

import (
	"strconv"
)

func BuildGreeting(name string, count int) string {
	return "Hello, " + name + "! You have " + strconv.Itoa(count) + " new messages."
}

// END
