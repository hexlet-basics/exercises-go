package solution

import (
	"strconv"
)

// BEGIN

func BuildGreeting(name string, count int) string {
	return "Hello, " + name + "! You have " + strconv.Itoa(count) + " new messages."
}

// END
